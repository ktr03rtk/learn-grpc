package client

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"reversi/build"
	"reversi/game"
	"reversi/gen/pb"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Reversi struct {
	sync.RWMutex
	started  bool
	finished bool
	me       *game.Player
	room     *game.Room
	game     *game.Game
}

func NewReversi() *Reversi {
	return &Reversi{}
}

func (r *Reversi) Run() int {
	if err := r.run(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func (r *Reversi) run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return errors.Wrapf(err, "failed to connect toserver")
	}
	defer conn.Close()

	if err := r.matching(ctx, pb.NewMatchingServiceClient(conn)); err != nil {
		return err
	}

	r.game = game.NewGame(r.me.Color)

	return r.play(ctx, pb.NewGameServiceClient(conn))

}

func (r *Reversi) matching(ctx context.Context, cli pb.MatchingServiceClient) error {
	stream, err := cli.JoinRoom(ctx, &pb.JoinRoomRequest{})
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	fmt.Println("requested matching")

	for {
		resp, err := stream.Recv()
		if err != nil {
			return err
		}

		if resp.GetStatus() == pb.JoinRoomResponse_MATCHED {
			r.room = build.Room(resp.GetRoom())
			r.me = build.Player(resp.GetMe())
			fmt.Printf("matched room_id=%d\n", resp.GetMe().GetId())

			return nil
		} else if resp.GetStatus() == pb.JoinRoomResponse_WAITING {
			fmt.Println("waiting for matching")
		}
	}
}

func (r *Reversi) play(ctx context.Context, cli pb.GameServiceClient) error {
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := cli.Play(c)
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	go func() {
		err := r.send(c, stream)
		if err != nil {
			cancel()
		}
	}()

	if err := r.receive(c, stream); err != nil {
		cancel()
		return err
	}

	return nil
}

func (r *Reversi) send(ctx context.Context, stream pb.GameService_PlayClient) error {
	for {
		r.RLock()

		if r.finished {
			r.RUnlock()
			return nil
		} else if !r.started {
			err := stream.Send(&pb.PlayRequest{
				RoomId: r.room.ID,
				Player: build.PBPlayer(r.me),
				Action: &pb.PlayRequest_Start{
					Start: &pb.PlayRequest_StartAction{},
				},
			})
			r.RUnlock()
			if err != nil {
				return err
			}

			for {
				r.RLock()
				if r.started {
					r.RUnlock()
					fmt.Println("ready go")
					break
				}
				r.RUnlock()
				fmt.Println("waiting until opponent player ready")
				time.Sleep(1 * time.Second)
			}
		} else {
			r.RUnlock()

			fmt.Print("input yout move(ex. A-1):")
			stdin := bufio.NewScanner(os.Stdin)
			stdin.Scan()

			text := stdin.Text()
			x, y, err := parseInput(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			r.Lock()
			_, err = r.game.Move(x, y, r.me.Color)
			r.Unlock()
			if err != nil {
				fmt.Println(err)
				continue
			}

			go func() {
				err = stream.Send(&pb.PlayRequest{
					RoomId: r.room.ID,
					Player: build.PBPlayer(r.me),
					Action: &pb.PlayRequest_Move{
						Move: &pb.PlayRequest_MoveAction{
							Move: &pb.Move{
								X: x,
								Y: y,
							},
						},
					},
				})
				if err != nil {
					fmt.Println(err)
				}
			}()

			ch := make(chan int)
			go func(ch chan int) {
				fmt.Println("")
				for i := 0; i < 5; i++ {
					fmt.Printf("freezing in %d second.\n", (5 - i))
					time.Sleep(1 * time.Second)

				}
				fmt.Println("")
				ch <- 0
			}(ch)
			<-ch

			select {
			case <-ctx.Done():
				return nil
			default:
			}
		}
	}
}

func parseInput(txt string) (int32, int32, error) {
	ss := strings.Split(txt, "-")
	if len(ss) != 2 {
		return 0, 0, fmt.Errorf("invalid input")
	}

	xs := ss[0]
	xrs := []rune(strings.ToUpper(xs))
	x := int32(xrs[0]-rune('A')) + 1
	if x < 1 || 8 < x {
		return 0, 0, fmt.Errorf("invalid input")
	}

	ys := ss[1]
	y, err := strconv.ParseInt(ys, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input")
	}
	if y < 1 || 8 < y {
		return 0, 0, fmt.Errorf("invalid input")
	}

	return x, int32(y), nil
}

func (r *Reversi) receive(ctx context.Context, stream pb.GameService_PlayClient) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			return err
		}

		r.Lock()
		switch res.GetEvent().(type) {
		case *pb.PlayResponse_Waiting:
		case *pb.PlayResponse_Ready:
			r.started = true
			r.game.Display(r.me.Color)
		case *pb.PlayResponse_Move:
			color := build.Color(res.GetMove().GetPlayer().GetColor())
			if color != r.me.Color {
				move := res.GetMove().GetMove()
				r.game.Move(move.GetX(), move.GetY(), color)
				fmt.Print("input your move:")
			}
		case *pb.PlayResponse_Finished:
			r.finished = true
			winner := build.Color(res.GetFinished().Winner)
			fmt.Println("")
			if winner == game.None {
				fmt.Println("Draw!")
			} else if winner == r.me.Color {
				fmt.Println("You Win!!")
			} else {
				fmt.Println("You Lose!")
			}

			r.Unlock()
			return nil
		}
		r.Unlock()

		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}
}

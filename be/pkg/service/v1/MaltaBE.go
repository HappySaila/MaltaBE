package v1

import (
	"context"
	"database/sql"
	v1 "github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	"log"
	"strings"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type maltaBEServer struct {
	db DB
	p Poker
	t AuthToken
}

func NewMaltaServiceServer(db *sql.DB) v1.MaltaBEServer {
	return &maltaBEServer{
		db: DB{db: db},
		p: Poker{},
		t: AuthToken{},
	}
}

func (s *maltaBEServer) Calculate(ctx context.Context, req *v1.CalculateRequest) (*v1.CalculateResponse, error) {
	//Authenticate
	if isAuthenticated, res := s.authenticate(req.Token); !isAuthenticated {
		return res, nil
	}

	s.p.s.Reset()

	log.Print("Calculating...")

	game := strings.Split(req.HackData, "\n")
	for _, hand := range game {
		if len(hand) == 0{
			break
		}
		s.p.compareHands(string(hand[:14]), string(hand[15:]))
	}

	log.Print("Writing to db...")
	data := s.p.s.GetData()
	res, err := s.db.AddData(ctx, s.p.s.DataForDB())
	if !res || err!=nil {
		log.Print("Error writing to DB")
	}

	if res {
		log.Print("Successfully wrote to remote DB")
	}

	return &v1.CalculateResponse{
		Api:   apiVersion,
		Data:  data,
	}, nil
}

func (s *maltaBEServer) authenticate(token string) (bool, *v1.CalculateResponse) {
	claims, err := s.t.Decode(token)
	if err != nil {
		return false, &v1.CalculateResponse{
			Api:   apiVersion,
			Data:  "User is not Authenticated and cannot execute this action.",
		}
	}

	if claims.Issuer != "HappySaila" {
		return false, &v1.CalculateResponse{
			Api:   apiVersion,
			Data:  "User is not Authenticated and cannot execute this action.",
		}
	}

	return true, nil
}

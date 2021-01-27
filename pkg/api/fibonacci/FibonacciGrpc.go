package fibonacci

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/outZoNe/go-fibonacci-service/lib/fibonacci"
	rs "github.com/outZoNe/go-fibonacci-service/lib/redis"
	fibonacciGrpc "github.com/outZoNe/go-fibonacci-service/pkg/api/proto"
	"log"
	"strconv"
)

var ctx = context.Background()

type GrpcFibonacci struct{}

func (s *GrpcFibonacci) GetFibonacci(ctx context.Context, req *fibonacciGrpc.FibonacciRequest) (*fibonacciGrpc.FibonacciResponse, error) {
	return &fibonacciGrpc.FibonacciResponse{Res: getResponseJson(req.GetStartNum(), req.GetEndNum())}, nil
}

func (s *GrpcFibonacci) Echo(ctx context.Context, req *fibonacciGrpc.FibonacciRequest) (*fibonacciGrpc.FibonacciResponse, error) {
	return &fibonacciGrpc.FibonacciResponse{Res: getResponseJson(req.GetStartNum(), req.GetEndNum())}, nil
}

func getResponseJson(x int32, y int32) string {
	if x > y {
		x, y = y, x
	}

	fibiArr := make([]fibonacci.JsonFibElem, y-x+1)
	var count = 0
	for x <= y {
		// Проверяем есть ли число в редисе, если есть, то вернем его, если нет, то будем вычислять его
		val, err := rs.Client().Get(ctx, strconv.Itoa(int(x))).Result()
		if val != "" && err != redis.Nil && err != nil {
			fibVal, _ := strconv.Atoi(val)
			fibiArr[count] = fibonacci.JsonFibElem{SerialNumber: int(x), Value: fibVal}
		} else {
			var fibVal = fibonacci.GetSerialFibNum(int(x))
			fibiArr[count] = fibonacci.JsonFibElem{SerialNumber: int(x), Value: fibVal}
			err := rs.Client().Set(ctx, strconv.Itoa(int(x)), strconv.Itoa(fibVal), 0).Err()
			if err != redis.Nil && err != nil {
				panic(err)
			}
		}
		x++
		count++
	}
	jsonString, err := json.Marshal(fibiArr)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonString)
}

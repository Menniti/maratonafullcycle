package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"simulator/entity"
	"simulator/queue"
	"strings"
	"time"
	"github.com/joho/godotenv"
	"testing"
)

func TestDestinationToJson(t *testing.T){
	order := entity.Order {
		Uuid: "uuid-generic"
	}

	expected := []byte(`{"order":"uuid-generic","lat":"9999","lng":"9999"}`)

	result := destinationToJson(order, "9999", "9999")

	isEqual := bytes.Compare(result, expected)

	if (isEqual != 0){
		t.Errorf("Expected %v, returned %v", string(expected), string(result))
	}
}
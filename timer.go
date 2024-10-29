package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	println("Starting timer...", time.Now().Format(time.RFC3339))
	println("Enter countdown time in format 'HH:MM:SS' (example: 01:30:00 for 1 hour 30 minutes):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	timeStr := strings.TrimSpace(input)
	if timeStr == "" {
		fmt.Println("Deadline cannot be empty")
		os.Exit(1)
	}
	var hours, minutes, seconds int
	_, err := fmt.Sscanf(timeStr, "%d:%d:%d", &hours, &minutes, &seconds)
	if err != nil {
		fmt.Println("Invalid time format")
		os.Exit(1)
	}
	v := time.Now().Add(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second)
	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(v)
		if timeRemaining.t <= 0 {
			fmt.Println("Time's up!")
			os.Exit(0)
		}
		fmt.Printf("Days: %02d, Hours: %02d, Minutes: %02d, Seconds: %02d\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)
	}
	fmt.Println(v)
}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	timeRemaining := t.Sub(currentTime)
	total := int(timeRemaining.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int((total % (60 * 60 * 24)) / (60 * 60))
	minutes := int((total % (60 * 60)) / 60)
	seconds := int(total % 60)
	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}

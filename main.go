package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	ssh "github.com/helloyi/go-sshclient"
	"github.com/joho/godotenv"
)

var (
	lg = fmt.Println
	lf = fmt.Printf
)

type Config struct {
	Url      string
	Name     string
	Password string
	Key      string
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := &Config{
		Url:      os.Getenv("URL"),
		Name:     os.Getenv("NAME"),
		Password: os.Getenv("PASSWORD"),
		Key:      os.Getenv("KEY"),
	}

	lf("%+v\n", cfg)

	sshy, err := NewSSHy(cfg)
	if err != nil {
		log.Fatal(err)
	}

	banner, err := sshy.getBanner()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		lf(banner)
		// Read input from the user
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}
		result, err := sshy.ExecOne(input)
		if err != nil {
			log.Fatal(err)
		}
		lg(result)

	}

}

type SSHy struct {
	Client *ssh.Client
}

func NewSSHy(cfg *Config) (*SSHy, error) {
	client, err := ssh.DialWithPasswd(cfg.Url, cfg.Name, cfg.Password)
	if err != nil {
		lg("1")
		return nil, err
	}
	return &SSHy{
		Client: client,
	}, nil
}

func (s *SSHy) ExecOne(cmd string) (string, error) {
	out, err := s.Client.Cmd(cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (s *SSHy) getBanner() (string, error) {
	whoami, err := s.ExecOne("whoami")
	if err != nil {
		return "", err
	}
	hostname, err := s.ExecOne("hostname")
	if err != nil {
		return "", err
	}
	pwd, err := s.ExecOne("pwd")
	if err != nil {
		return "", err
	}
	banner := fmt.Sprintf("%s@%s:%s$ ",
		strings.TrimSuffix(whoami, "\n"),
		strings.TrimSuffix(hostname, "\n"),
		strings.TrimSuffix(pwd, "\n"))
	return banner, nil
}

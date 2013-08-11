package main

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/andybons/hipchat"
	"github.com/google/go-github/github"
	"os"
)

type Project struct {
	IO            ProjectIO
	Conf          *Config
	GithubOrg     string
	GithubRepo    string
	HipchatRoom   string
	Error         bool
	HipchatClient *hipchat.Client
	GithubClient  *github.Client
	HipchatRooms  []hipchat.Room
	GithubRepos   []github.Repository
}

func (p *Project) ReadConfig() {
	config, err := ReadConfig()
	if err != nil {
		p.IO.say("Config file not found. Please run `hygo init`")
		os.Exit(1)
	}
	p.Conf = config
}

func (p *Project) PromptForConfig() {
	config := Config{
		GithubAccessToken: p.IO.ask("Enter Github Access Token:"),
		HipchatAuthToken:  p.IO.ask("Enter Hipchat Auth Token:"),
	}
	p.Conf = &config
}

func (p *Project) WriteConfig() {
	p.Conf.Write()
}

func (p *Project) InitHipchatClient() {
	p.HipchatClient = &hipchat.Client{AuthToken: p.Conf.HipchatAuthToken}
}

func (p *Project) PromptForHipchatRoom() {
	p.HipchatRoom = p.IO.ask("Enter Hipchat Room Name:")
}

func (p *Project) GetHipchatRooms() {
	rooms, err := p.HipchatClient.RoomList()

	if err != nil {
		println("Hipchat Error:")
		println(err.Error())
		os.Exit(1)
	}
	p.HipchatRooms = rooms
}

func (p *Project) ListHipchatRooms() {
	p.ReadConfig()
	p.InitHipchatClient()
	p.GetHipchatRooms()
	for _, room := range p.HipchatRooms {
		println(room.Name)
	}
}

func (p *Project) InitGithubClient() {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: p.Conf.GithubAccessToken},
	}
	p.GithubClient = github.NewClient(t.Client())
}

func (p *Project) PromptForGithubOrg() {
	p.GithubOrg = p.IO.ask("Enter Github Org:")
}

func (p *Project) PromptForGithubRepo() {
	p.GithubRepo = p.IO.ask("Enter Github Repo:")
}

func (p *Project) GetGithubRepos() {
	repos, _, _ := p.GithubClient.Repositories.ListByOrg(p.GithubOrg, nil)
	p.GithubRepos = repos
}

func (p *Project) ListGithubRepos() {
	p.ReadConfig()
	p.InitGithubClient()
	p.PromptForGithubOrg()
	p.GetGithubRepos()
	for _, repo := range p.GithubRepos {
		p.IO.say(repo.Name)
	}
}

func (p *Project) AddGithubHipchatHook() {
	p.ReadConfig()
	p.InitHipchatClient()
	p.InitGithubClient()
	p.PromptForGithubOrg()
	p.PromptForGithubRepo()
	p.PromptForHipchatRoom()
	p.DoAddGithubHipchatHook()
}

func (p *Project) DoAddGithubHipchatHook() {

	hookConf := map[string]interface{}{
		"room":       p.HipchatRoom,
		"auth_token": p.Conf.HipchatAuthToken,
	}
	hook := &github.Hook{
		Name:   "hipchat",
		Config: hookConf,
	}
	_, _, err := p.GithubClient.Repositories.CreateHook(p.GithubOrg, p.GithubRepo, hook)
	if err != nil {
		p.IO.say(err.Error())
		os.Exit(1)
	}
	p.IO.say("Hook added!")
}

package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Acitvity struct {
	Type      string `json: "type"`
	Repo      Repo   `json: "repo"`
	CreatedAt string `json: "created_at"`
	Payload   struct {
		Action  string `json:"action"`
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"mesage"`
		} `json:"commits"`
	} `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

func GetActivity(name string) ([]Acitvity, error) {
	var acts []Acitvity
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", name))
	if err != nil {
		return acts, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return acts, fmt.Errorf("user not found")
	}

	if resp.StatusCode != http.StatusOK {
		return acts, fmt.Errorf("error fetching data: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&acts)
	if err != nil {
		return acts, fmt.Errorf("error reading response body")
	}

	return acts, nil
}

func FilterActivities(t string, acts []Acitvity) ([]Acitvity, error) {
	res := make([]Acitvity, 0)
	if len(acts) == 0 {
		return res, fmt.Errorf("no activity found")
	}

	tp := ""

	switch t {
	case "push":
		tp = "PushEvent"
	case "star":
		tp = "WatchEvent"
	case "issue":
		tp = "IssuesEvent"
	case "fork":
		tp = "ForkEvent"
	case "create":
		tp = "CreateEvent"
	case "delete":
		tp = "DeleteEvent"
	}

	for _, act := range acts {
		if act.Type == tp {
			res = append(res, act)
		}
	}

	return res, nil
}

func PrintActivities(acts []Acitvity) {
	fmt.Println("Output: ")
	for _, act := range acts {
		switch act.Type {
		case "PushEvent":
			count := len(act.Payload.Commits)
			fmt.Printf("- Pushed %d commits to %s\n", count, act.Repo.Name)
		case "IssuesEvent":
			action := ""
			if act.Payload.Action == "opened" {
				action = "Opened"
			} else {
				action = "Closed"
			}
			fmt.Printf("- %s an issue in %s\n", action, act.Repo.Name)
		case "WatchEvent":
			action := ""
			if act.Payload.Action == "started" {
				action = "Starred"
			} else {
				action = "Unstarred"
			}
			fmt.Printf("- %s %s\n", action, act.Repo.Name)
		case "ForkEvent":
			fmt.Printf("- Forked %s\n", act.Repo.Name)
		case "CreateEvent":
			fmt.Printf("- Created %s in %s\n", act.Payload.RefType, act.Repo.Name)
		case "DeleteEvent":
			fmt.Printf("- Deleted %s %s in %s\n", act.Payload.RefType, act.Payload.Ref, act.Repo.Name)
		default:
			fmt.Printf("- %s in %s\n", act.Type, act.Repo.Name)
		}
	}
}

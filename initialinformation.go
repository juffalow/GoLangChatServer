package main

type InitialInformation struct {
    Usernames []string `json:"usernames"`
    Messages []Message `json:"messages"`
}

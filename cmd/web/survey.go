package main

import (
	"fmt"
)

type Survey struct {
	Id        int
	Name      string
	Questions []Question
	Current   int
}

type Question struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Options     []Option `json:"options"`
	Counter     int      `json:"counter"`
	Answers     Answers  `json:"answers"`
}

type Answers struct {
	IpList map[string]*Option
}

type Option struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	//IpList []string `json:"ip_list"`
}

type Result struct {
	Id       int
	question *Question
	result   int
}

type ResponsePayload struct {
	Id int
}

type NewQuestionPayload struct {
	Name string
	Q1   string
	Q2   string
	Q3   string
	Pin  string
}

func CreateDummySurvey() Survey {
	return Survey{
		Id:   1,
		Name: "BTS Опрос!",
		Questions: []Question{
			{
				Id:          1,
				Description: "Что думаете о заявленной теме?",
				Options: []Option{
					{Name: "Просто интересно"},
					{Name: "Это путь в никуда"},
					{Name: "Это светлое будущее"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          2,
				Description: "Пользуетесь облаками?",
				Options: []Option{
					{Name: "Постоянно"},
					{Name: "Пару раз"},
					{Name: "Никогда"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          3,
				Description: "С какими моделями сервисов сталкивались?",
				Options: []Option{
					{Name: "SaaS"},
					{Name: "PaaS"},
					{Name: "IaaS"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          4,
				Description: "Как вы думаете, насколько часто вы пользуетесь ПО на микросервисах?",
				Options: []Option{
					{Name: "Никогда"},
					{Name: "Все время"},
					{Name: "Пару раз в год"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          5,
				Description: "Можно запустить людей в космос по скраму?",
				Options: []Option{
					{Name: "Конечно нет, у них даже двигателей своих нет"},
					{Name: "Уже запустили"},
					{Name: "О чем речь?"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          6,
				Description: "Сможете прожить без FOSS?",
				Options: []Option{
					{Name: "Предки выживали, и мы выживем"},
					{Name: "Пути назад нет"},
					{Name: "А зачем?"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          7,
				Description: "Хотите стать элитой?",
				Options: []Option{
					{Name: "Уже стали"},
					{Name: "Конечно"},
					{Name: "Не мое"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          8,
				Description: "Вселенная стала понятнее?",
				Options: []Option{
					{Name: "Намного"},
					{Name: "Слегка"},
					{Name: "Совсем нет"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
			{
				Id:          9,
				Description: "О чем интересно послушать в следующий раз?",
				Options: []Option{
					{Name: "Про ифнраструктуру"},
					{Name: "Про разработку"},
					{Name: "Про будущее"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
		},
		Current: 0,
	}
}

func (survey *Survey) AddQuestion(name, q1, q2, q3 string) {
	question := Question{
		Id:          len((*survey).Questions),
		Description: name,
		Options: []Option{
			{Name: q1},
			{Name: q2},
			{Name: q3},
		},
		Answers: Answers{IpList: map[string]*Option{}},
	}

	(*survey).Questions = append((*survey).Questions, question)
	(*survey).Current = len((*survey).Questions) - 1
	fmt.Printf("survey = %v\n", *survey)
}

func (survey *Survey) SetQuestion(pos int) {
	(*survey).Current = pos
}

func (survey *Survey) Increment(i int, ip string) {
	lastQuestion := (*survey).CurrentQuestion()
	if _, ok := lastQuestion.Answers.IpList[ip]; !ok {
		lastQuestion.Answers.IpList[ip] = &lastQuestion.Options[i]
		lastQuestion.Answers.IpList[ip].Count++
	} else {
		lastQuestion.Answers.IpList[ip].Count--
		lastQuestion.Answers.IpList[ip] = &lastQuestion.Options[i]
		lastQuestion.Answers.IpList[ip].Count++
	}
}

func (survey *Survey) CurrentQuestion() *Question {
	return &(*survey).Questions[(*survey).Current]
}

package telegram

import (
	"algoBot/internal/db/repository"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	bigOTopic          = "О-большое"
	arrayTopic         = "Массив"
	linkedListTopic    = "Связный список"
	stackTopic         = "Стек"
	hashMapTable       = "Хеш-таблица"
	bTreeTopic         = "Бинарное дерево и Бинарное дерево поиска"
	bSearchTopic       = "Бинарный поиск"
	twoPointersTopic   = "Метод двух указателей"
	slidingWindowTopic = "Метод Скользящего окна"
	BFSTopic           = "BFS(Поиск в ширину)"
	DFSTopic           = "DFS(Поиск в глубину)"
	backtrackTopic     = "Поиск с возвратом"
	queueTopic         = "Очередь"
	recursionTopic     = "Рекурсия"
	sortsTopic         = "Сортировки"
)

const (
	cmdStart         = "start"
	cmdHelp          = "help"
	cmdGuide         = "guide"
	cmdBigO          = "bigO"
	cmdArray         = "array"
	cmdLinkedList    = "linkedList"
	cmdStack         = "stack"
	cmdHashTable     = "hashTable"
	cmdBTree         = "binaryTree"
	cmdBSearch       = "binarySearch"
	cmdTwoPointers   = "twoPointers"
	cmdSlidingWindow = "slidingWindow"
	cmdBFS           = "BFS"
	cmdDFS           = "DFS"
	cmdBacktrack     = "backtrack"
	cmdQueue         = "queue"
	cmdRecursion     = "recursion"
	cmdSorts         = "sorts"
)

func (b *Bot) handleCmd(message *tgbotapi.Message) error {
	switch message.Command() {
	case cmdStart:
		return b.handleMsgCmd(message, msgStart)
	case cmdHelp:
		return b.handleMsgCmd(message, msgHelp)
	case cmdGuide:
		return b.handleMsgCmd(message, msgGuide)
	case cmdBigO:
		return b.handleTopic(message, bigOTopic)
	case cmdArray:
		return b.handleTopic(message, arrayTopic)
	case cmdLinkedList:
		return b.handleTopic(message, linkedListTopic)
	case cmdStack:
		return b.handleTopic(message, stackTopic)
	case cmdHashTable:
		return b.handleTopic(message, hashMapTable)
	case cmdBTree:
		return b.handleTopic(message, bTreeTopic)
	case cmdBSearch:
		return b.handleTopic(message, bSearchTopic)
	case cmdTwoPointers:
		return b.handleTopic(message, twoPointersTopic)
	case cmdSlidingWindow:
		return b.handleTopic(message, slidingWindowTopic)
	case cmdBFS:
		return b.handleTopic(message, BFSTopic)
	case cmdDFS:
		return b.handleTopic(message, DFSTopic)
	case cmdBacktrack:
		return b.handleTopic(message, backtrackTopic)
	case cmdQueue:
		return b.handleTopic(message, queueTopic)
	case cmdRecursion:
		return b.handleTopic(message, recursionTopic)
	case cmdSorts:
		return b.handleTopic(message, sortsTopic)
	default:
		return b.handleMsg(message)
	}
}

func (b *Bot) handleMsgCmd(message *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleTopic(message *tgbotapi.Message, topic string) error {
	page, err := b.repository.Get(context.Background(), topic)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, format(page))
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleMsg(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, msgUnknown)
	_, err := b.bot.Send(msg)
	return err
}

func format(page *repository.Page) string {
	return fmt.Sprintf("%s\n\n%s\n\nПолезные материалы: %s", page.Topic, page.ShortDesc, page.UsefulMaterials)
}

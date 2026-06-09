package task

import (
	"context"
	"errors"
	"time"

	"github.com/sevelfatt/taskverse-backend/lib"
	"github.com/sevelfatt/taskverse-backend/models"
	"github.com/sevelfatt/taskverse-backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetTaskByUUIDService(taskUUID string) (any, error){
	db := lib.MongoClient.Database("taskverse")

	taskCollection := db.Collection("tasks")

	var task any
	err := taskCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "task.uuid", Value: taskUUID}}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func DeleteTaskByUUIDService(taskUUID string) error {
	db := lib.MongoClient.Database("taskverse")

	taskCollection := db.Collection("tasks")

	_, err := taskCollection.DeleteOne(context.TODO(), bson.D{bson.E{Key: "task.uuid", Value: taskUUID}})
	if err != nil {
		return err
	}

	return nil
}

func GetAllTasksByUserUUIDService(userUUID string) ([]any, error) {
	db := lib.MongoClient.Database("taskverse")

	taskCollection := db.Collection("tasks")

	var tasks []any
	cursor, err := taskCollection.Find(context.TODO(), bson.D{bson.E{Key: "task.useruuid", Value: userUUID}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTaskService(userUUID string, title string, taskType string, daysInWeek []time.Weekday, startDate time.Time, endDate time.Time) (any, error) {
	switch taskType {
	case "habit":
		return createHabitTask(models.HabitTask{
			Task: models.Task{
				UUID: utils.GenerateUUID(),
				UserUUID: userUUID,
				Title: title,
				Type: taskType,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			DaysInWeek: daysInWeek,
		})
	case "one_time":
		return createOneTimeTask(models.OneTimeTask{
			Task: models.Task{
				UUID: utils.GenerateUUID(),
				UserUUID: userUUID,
				Title: title,
				Type: taskType,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Date: startDate,
		})
	case "weekly":
		return createWeeklyMonthlyTask(models.WeeklyAndMonthyTask{
			Task: models.Task{
				UUID: utils.GenerateUUID(),
				UserUUID: userUUID,
				Title: title,
				Type: taskType,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			StartDate: startDate,
			EndDate: endDate,
		})
	case "monthly":
		return createWeeklyMonthlyTask(models.WeeklyAndMonthyTask{
			Task: models.Task{
				UUID: utils.GenerateUUID(),
				UserUUID: userUUID,
				Title: title,
				Type: taskType,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			StartDate: startDate,
			EndDate: endDate,
		})
	default:
		return nil, errors.New("Invalid task type")
	}
}

func createWeeklyMonthlyTask(task models.WeeklyAndMonthyTask) (models.WeeklyAndMonthyTask, error){
	db := lib.MongoClient.Database("taskverse")

	taskCollection := db.Collection("tasks")

	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.WeeklyAndMonthyTask{}, err
	}

	return task, nil
}

func createOneTimeTask(task models.OneTimeTask) (models.OneTimeTask, error){
	db := lib.MongoClient.Database("taskverse")

	taskCollection := db.Collection("tasks")

	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.OneTimeTask{}, err
	}

	return task, nil
}

func createHabitTask(task models.HabitTask) (models.HabitTask, error){
	db := lib.MongoClient.Database("taskverse")

	taskCollection := db.Collection("tasks")

	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.HabitTask{}, err
	}

	return task, nil
}

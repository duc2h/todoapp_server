package biz

import "github.com/hoangduc02011998/todo_server/models"

func GetAllTask() *models.ResponseModel {
	tasks, err := taskRepo.FindAll()

	if err != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		}
	}

	return &models.ResponseModel{
		Status:  models.APIStatus.Ok,
		Message: models.APIStatus.Ok,
		Data:    tasks,
	}
}

func GetTaskByName(taskName string) *models.ResponseModel {
	task, _ := taskRepo.FindTaskByName(taskName)

	if task == nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.NotFound,
			Message: models.APIStatus.NotFound,
		}
	}

	return &models.ResponseModel{
		Status:  models.APIStatus.Ok,
		Message: models.APIStatus.Ok,
		Data:    task,
	}
}

func CreateTask(input models.Task) *models.ResponseModel {
	task, _ := taskRepo.FindTaskByName(input.TaskName)

	if task != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Existed,
			Message: "Task is available",
		}
	}

	err := taskRepo.Insert(input)
	if err != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		}
	}

	return &models.ResponseModel{
		Status:  models.APIStatus.Ok,
		Message: models.APIStatus.Ok,
	}
}

func UpdateTask(name string, status bool) *models.ResponseModel {
	err := taskRepo.Update(name, status)
	if err != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		}
	}

	return &models.ResponseModel{
		Status:  models.APIStatus.Ok,
		Message: models.APIStatus.Ok,
	}
}

func DeleteTask(name string) *models.ResponseModel {
	err := taskRepo.Delete(name)
	if err != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		}
	}

	return &models.ResponseModel{
		Status:  models.APIStatus.Ok,
		Message: models.APIStatus.Ok,
	}
}

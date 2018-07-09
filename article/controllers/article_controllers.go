package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/daniel/basic_microservice/article/connections"
	"github.com/daniel/basic_microservice/article/models"
	"github.com/daniel/basic_microservice/article/commons/response"
	"strconv"
)

func Index(context echo.Context) error {
	db, err := connections.GetConnection()

	if err != nil{
		return dispathError(context)
	}

	articles := make([]models.Article, 0)

	db.Find(&articles)

	if len(articles) == 0{
		return response.DisplayMessage(
			http.StatusNotFound,
			"No se encontraron registros en la base de datos",
			"no-content",
			context,
		)
	}

	return response.DisplayMessage(
		http.StatusOK,
		"La solicitud fue gestionada exitosamente",
		articles,
		context,
	)
}

func Show(context echo.Context) error{

	var article models.Article
	db, err := connections.GetConnection()

	if err != nil{
		return dispathError(context)
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil{
		return response.DisplayMessage(
			http.StatusBadRequest,
			"La peticion no envio los parametros correctos para obtener el registro deseado",
			"no-content",
			context,
		)
	}

	db.Find(&article, models.Article{ID: id })

	if article.ID == 0 {
		return response.DisplayMessage(
			http.StatusNotFound,
			"No se encontro el registro",
			"no-content",
			context,
		)
	}

	return response.DisplayMessage(
		http.StatusOK,
		"La solicitud fue gestionada exitosamente",
		article,
		context,
	)

}

func Create(context echo.Context) error{

	article := new(models.Article)
	db, err := connections.GetConnection()

	if err != nil {
		return dispathError(context)
	}

	if err = context.Bind(article); err != nil {
		return response.DisplayMessage (
			http.StatusBadRequest,
			"La peticion no envio los parametros correctos para registrar el articulo",
			"no-content",
			context,
		)

	}

	err = db.Create(&article).Error

	if err != nil {
		return response.DisplayMessage(
			http.StatusBadRequest,
			"Ha ocurrido un error al momento de registrar el articulo \n Tal vez no enviaste los datos importantes o el articulo que se desea registrar ya existia",
			"no-content",
			context,
		)
	}

	return response.DisplayMessage(
		http.StatusCreated,
		"Se ha registrado exitosamente el articulo",
		"no-content",
		context,
	)

	return  context.String(http.StatusOK, "Crear")
}

func Update(context echo.Context) error{

	article := new(models.Article)

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return response.DisplayMessage(
			http.StatusBadRequest,
			"La peticion no envio los parametros correctos para actualizar el registro deseado",
			"no-content",
			context,
		)
	}

	db, err := connections.GetConnection()

	if err != nil {
		return dispathError(context)
	}

	if err = context.Bind(article); err != nil{
		return response.DisplayMessage(
			http.StatusBadRequest,
			"La peticion no envio los parametros correctos para actualizar el articulo",
			"no-content",
			context,
		)
	}

	article.ID = id

	err = db.Model(&article).Where("id = ?", article.ID).Updates(map[string]interface{}{
		"title": article.Title,
		"content": article.Content,
		"description": article.Description,
		"image": article.Image,
	}).Error

	if err != nil {
		return response.DisplayMessage(
			http.StatusBadRequest,
			"La peticion no envio los parametros correctos para actualizar el articulo",
			"no-content",
			context,
		)
	}

	return response.DisplayMessage(
		http.StatusCreated,
		"El registro fue actualizado exitosamente",
		"no-content",
		context,
	)

}

func Delete(context echo.Context) error{
	var id int

	db, err := connections.GetConnection()

	if err != nil {
		dispathError(context)
	}

	id, err =  strconv.Atoi(context.Param("id"))
	if err != nil {
		return response.DisplayMessage(
			http.StatusBadRequest,
			"La peticion no envio los parametros correctos para eliminar el registro deseado",
			"no-content",
			context,
		)
	}


	err = db.Delete(models.Article{}, "id = ?", id).Error

	if err != nil {
		return response.DisplayMessage(
			http.StatusBadRequest,
			"No se pudo eliminar el registro deseado",
			"no-content",
			context,
		)
	}

	return response.DisplayMessage(
		http.StatusCreated,
		"El registro se ha eliminado exitosamente",
		"no-content",
		context,
	)
}

func dispathError(context echo.Context) error{
	return response.DisplayMessage(
		http.StatusInternalServerError,
		"Hubo un error en el servidor",
		"no-content",
		context,
	)
}
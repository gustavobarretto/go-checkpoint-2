package appointmenthandler

import (
	"checkpoint-2/internal/application/usecase"
	"checkpoint-2/internal/domain"
	"checkpoint-2/pkg/web"
	"errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type appointmentRequest struct {
	
}

type appointmentHandler struct {
	appointmentGroup	gin.RouterGroup
	appointmentService usecase.Appointment
}

func (a *appointmentHandler) ConfigureAppointmentRouter() {
	a.appointmentGroup.POST("", a.post)
	a.appointmentGroup.GET(":id", a.get)
	a.appointmentGroup.GET("", a.getAll)
	a.appointmentGroup.PUT(":id", a.put)
	a.appointmentGroup.PATCH(":id", a.patch)
	a.appointmentGroup.DELETE(":id", a.delete)
}

func (a *appointmentHandler) post(ctx *gin.Context) {
	var appointment domain.CreateAppointment
	err := ctx.ShouldBindJSON(&appointment); if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	err = a.appointmentService.Post(appointment); if err != nil {
		web.Failure(ctx, 500, err)
		return
	}
	ctx.Status(201)
}

func (a *appointmentHandler) get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		web.Failure(ctx, 400, errors.New("no id sent"))
		return
	}

	idConverted, err := strconv.Atoi(id); if err != nil {
		web.Failure(ctx, 400, errors.New("incorrect id sent. must be a number"))
		return
	}

	appointment, err := a.appointmentService.Get(idConverted); if err != nil {
		web.Failure(ctx, 500, errors.New("errors getting entity"))
		return
	}

	if reflect.DeepEqual(appointment, domain.Appointment{}) {
		web.Failure(ctx, 404, errors.New("entity not found"))
		return
	}

	ctx.JSON(200, appointment)
}

func (a *appointmentHandler) getAll(ctx *gin.Context) {
	appointments, err := a.appointmentService.GetAll(); if err != nil {
		web.Failure(ctx, 500, err)
		return
	}
	ctx.JSON(200, appointments)
}

func (a *appointmentHandler) put(ctx *gin.Context) {
	var appointment domain.UpdateAppointment
	id := ctx.Param("id")
	if id == "" {
		web.Failure(ctx, 400, errors.New("no id sent"))
		return
	}

	idConverted, err := strconv.Atoi(id); if err != nil {
		web.Failure(ctx, 400, errors.New("incorrect id sent. must be a number"))
		return
	}

	err = ctx.ShouldBindJSON(&appointment); if err != nil {
		web.Failure(ctx, 400, err)
		return
	}

	_, err = a.appointmentService.Get(idConverted); if err != nil {
		web.Failure(ctx, 500, errors.New("errors getting entity"))
		return
	}

	if reflect.DeepEqual(appointment, domain.UpdateAppointment{}) {
		web.Failure(ctx, 404, errors.New("entity not found"))
		return
	}

	err = a.appointmentService.Put(idConverted, domain.UpdateAppointment{
		appointment.PatientId,
		appointment.DentistId,
		appointment.AppointmentDate,
	}); if err != nil {
		web.Failure(ctx, 500, err)
		return
	}

	ctx.Status(204)
}
func (a *appointmentHandler) patch(ctx *gin.Context) {
	var appointment domain.PatchAppointment
	id := ctx.Param("id")
	if id == "" {
		web.Failure(ctx, 400, errors.New("no id sent"))
		return
	}

	idConverted, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, errors.New("incorrect id sent. must be a number"))
		return
	}

	err = ctx.ShouldBindJSON(&appointment)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}

	_, err = a.appointmentService.Get(idConverted)
	if err != nil {
		web.Failure(ctx, 500, errors.New("errors getting entity"))
		return
	}

	if reflect.DeepEqual(appointment, domain.PatchAppointment{}) {
		web.Failure(ctx, 404, errors.New("entity not found"))
		return
	}

	err = a.appointmentService.Patch(idConverted, domain.PatchAppointment{AppointmentDate: appointment.AppointmentDate})
	if err != nil {
		web.Failure(ctx, 500, err)
		return
	}

	ctx.Status(204)
}

func (a *appointmentHandler) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		web.Failure(ctx, 400, errors.New("no id sent"))
		return
	}

	idConverted, err := strconv.Atoi(id); if err != nil {
		web.Failure(ctx, 400, errors.New("incorrect id sent. must be a number"))
		return
	}

	appointment, err := a.appointmentService.Get(idConverted); if err != nil {
		web.Failure(ctx, 500, errors.New("errors getting entity"))
		return
	}

	if reflect.DeepEqual(appointment, domain.Dentist{}) {
		web.Failure(ctx, 404, errors.New("entity not found"))
		return
	}

	err = a.appointmentService.Delete(idConverted); if err != nil {
		web.Failure(ctx, 500, err)
		return
	}

	ctx.Status(204)

}

func NewAppointmentHandler(routerGroup *gin.RouterGroup, service usecase.Appointment) appointmentHandler {
	return appointmentHandler{*routerGroup, service}
}
package routers

import (
	"server/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api",

		// user login
		beego.NSRouter("/login",
			&controllers.UserController{},
			"post:Login",
		),

		// create leave request form (for all role except admin and director)
		beego.NSRouter("/leave/:id:int ",
			&controllers.LeaveController{},
			"post:PostLeaveRequest",
		),
		// create leave request form (for all role except admin and director)
		beego.NSRouter("/supervisor/leave/:id:int ",
			&controllers.LeaveController{},
			"post:PostLeaveRequestSupervisor",
		),

		// ========================= admin ========================= //
		// get leave request success
		beego.NSRouter("/admin/leave",
			&controllers.AdminController{},
			"get:GetRequestAccept",
		),
		// register user
		beego.NSRouter("/admin/user/register",
			&controllers.AdminController{},
			"post:CreateUser",
		),
		// get all user
		beego.NSRouter("/admin/user",
			&controllers.AdminController{},
			"get:GetUsers;post:CreateUser",
		),
		// get one user, update one user, delete one user
		beego.NSRouter("/admin/user/:id:int",
			&controllers.AdminController{},
			"get:GetUser;put:UpdateUser;delete:DeleteUser",
		),
		// delete user
		beego.NSRouter("/admin/user/delete/:id:int",
			&controllers.AdminController{},
			"delete:DeleteUser",
		),

		// ========================= employee ========================= //
		// get pending request in employee
		beego.NSRouter("/employee/pending/:id:int ",
			&controllers.UserController{},
			"get:GetRequestPending",
		),
		// get accept request in employee
		beego.NSRouter("/employee/accept/:id:int ",
			&controllers.UserController{},
			"get:GetRequestAccept",
		),
		// get reject request in employee
		beego.NSRouter("/employee/reject/:id:int ",
			&controllers.UserController{},
			"get:GetRequestReject",
		),

		// ========================= supervisor ========================= //
		// accept status by supervisor
		beego.NSRouter("/employee/accept/:id:int/:enumber:int",
			&controllers.UserController{},
			"put:AcceptStatusBySupervisor",
		),
		// reject status by supervisor
		beego.NSRouter("/employee/reject/:id:int/:enumber:int",
			&controllers.UserController{},
			"put:RejectStatusBySupervisor",
		),
		// get status pending in supervisor
		beego.NSRouter("/supervisor/pending/:id:int ",
			&controllers.UserController{},
			"get:GetPendingLeave",
		),
		// get status accept in supervisor
		beego.NSRouter("/supervisor/accept/:id:int ",
			&controllers.UserController{},
			"get:GetAcceptLeave",
		),
		// get status reject in supervisor
		beego.NSRouter("/supervisor/reject/:id:int ",
			&controllers.UserController{},
			"get:GetRejectLeave",
		),

		// ========================= director ========================= //
		// accept status by director
		beego.NSRouter("/director/accept/:id:int/:enumber:int",
			&controllers.DirectorController{},
			"put:AcceptStatusByDirector",
		),
		// reject status by director
		beego.NSRouter("/director/reject/:id:int/:enumber:int",
			&controllers.DirectorController{},
			"put:RejectStatusByDirector",
		),
		// get status pending in director
		beego.NSRouter("/director/pending/",
			&controllers.DirectorController{},
			"get:GetDirectorPendingLeave",
		),
		// get status accept in director
		beego.NSRouter("/director/accept/",
			&controllers.DirectorController{},
			"get:GetDirectorAcceptLeave",
		),
		// get status reject in director
		beego.NSRouter("/director/reject/",
			&controllers.DirectorController{},
			"get:GetDirectorRejectLeave",
		),
	)
	beego.AddNamespace(ns)
}

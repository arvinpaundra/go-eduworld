package routes

import (
	authCtrl "github.com/arvinpaundra/go-eduworld/internal/controllers/auth"
	courseCtrl "github.com/arvinpaundra/go-eduworld/internal/controllers/course"
	materialCtrl "github.com/arvinpaundra/go-eduworld/internal/controllers/material"
	moduleCtrl "github.com/arvinpaundra/go-eduworld/internal/controllers/module"
	authSrvc "github.com/arvinpaundra/go-eduworld/internal/services/auth"
	courseSrvc "github.com/arvinpaundra/go-eduworld/internal/services/course"
	materialSrvc "github.com/arvinpaundra/go-eduworld/internal/services/material"
	moduleSrvc "github.com/arvinpaundra/go-eduworld/internal/services/module"
	"github.com/uptrace/bun"

	"github.com/arvinpaundra/go-eduworld/internal/drivers"
	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Sql      *bun.DB
	Redis    *redis.Client
	Fiber    *fiber.App
	RabbitMQ *amqp091.Connection
}

func (c *Config) Start() {
	// dependecy injection
	deviceRepository := drivers.NewDeviceRepository(c.Sql)
	sessionRepository := drivers.NewSessionRepository(c.Sql)
	userRepository := drivers.NewUserRepository(c.Sql)
	interestRepository := drivers.NewInterestRepository(c.Sql)
	categoryRepository := drivers.NewCategoryRepository(c.Sql)
	courseRepository := drivers.NewCourseRepository(c.Sql)
	moduleRepository := drivers.NewModuleRepository(c.Sql)
	materialRepository := drivers.NewMaterialRepository(c.Sql)
	cacheRepository := drivers.NewCacheRepository(c.Redis)

	authService := authSrvc.NewAuthService(
		userRepository,
		deviceRepository,
		sessionRepository,
		interestRepository,
		cacheRepository,
	)
	courseService := courseSrvc.NewCourseService(
		userRepository,
		interestRepository,
		categoryRepository,
		courseRepository,
		moduleRepository,
		materialRepository,
		cacheRepository,
	)
	moduleService := moduleSrvc.NewModuleService(
		courseRepository,
		moduleRepository,
	)
	materialService := materialSrvc.NewMaterialService(
		courseRepository,
		moduleRepository,
		materialRepository,
	)

	authController := authCtrl.NewAuthController(authService)
	courseController := courseCtrl.NewCourseController(courseService)
	moduleController := moduleCtrl.NewModuleController(moduleService)
	materialController := materialCtrl.NewMaterialController(materialService)

	// group app version
	v1 := c.Fiber.Group("/api/v1")

	// auth routes
	auth := v1.Group("/auth")
	auth.Post("/login", authController.HandlerLogin)
	auth.Post("/register", authController.HandlerRegister)
	auth.Post("/logout", authController.HandlerLogout)

	// course routes
	course := v1.Group("/courses")
	course.Get("", courseController.HandlerFindAll)
	course.Post("", courseController.HandlerCreate)

	courseDetail := course.Group("/:course_id")
	courseDetail.Get("", courseController.HandlerFindDetail)
	courseDetail.Put("", courseController.HandlerUpdate)
	courseDetail.Delete("", courseController.HandlerRemove)

	// module routes
	module := courseDetail.Group("/modules")
	module.Post("", moduleController.HandlerCreate)

	moduleDetail := module.Group("/:module_id")
	moduleDetail.Put("", moduleController.HandlerUpdate)
	moduleDetail.Delete("", moduleController.HandlerRemove)

	// material routes
	material := moduleDetail.Group("/materials")
	material.Post("", materialController.HandlerCreate)

	materialDetail := material.Group("/:material_id")
	materialDetail.Get("", materialController.HandlerFindOne)
	materialDetail.Put("", materialController.HandlerUpdate)
	materialDetail.Delete("", materialController.HandlerRemove)

	// mentor routes
	mentor := v1.Group("/mentors")
	mentorDetail := mentor.Group("/:mentor_id")
	mentorDetail.Get("/courses", courseController.HandlerFindByMentor)
}

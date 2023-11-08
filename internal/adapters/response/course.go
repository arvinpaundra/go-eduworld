package response

import (
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
)

type (
	Course struct {
		ID                 string    `json:"id"`
		CategoryId         string    `json:"category_id"`
		MentorId           string    `json:"mentor_id"`
		InterestId         string    `json:"interest_id"`
		Title              string    `json:"title"`
		Category           string    `json:"category"`
		Interest           string    `json:"interest"`
		Mentor             string    `json:"mentor"`
		Level              string    `json:"level"`
		Type               string    `json:"type"`
		IsPublished        bool      `json:"is_published"`
		Price              *int      `json:"price"`
		Thumbnail          *string   `json:"thumbnail"`
		CategoryImage      *string   `json:"category_image"`
		MentorProfileImage *string   `json:"mentor_profile_image"`
		CreatedAt          string    `json:"created_at"`
		UpdatedAt          string    `json:"updated_at"`
		Modules            []*Module `json:"modules"`
	}

	Module struct {
		ID          string      `json:"id"`
		CourseId    string      `json:"course_id"`
		Title       string      `json:"title"`
		Description *string     `json:"description"`
		Materials   []*Material `json:"materials"`
	}

	Material struct {
		ID          string  `json:"id"`
		CourseId    string  `json:"course_id"`
		ModuleId    string  `json:"module_id"`
		Title       string  `json:"title"`
		Url         string  `json:"url"`
		Description *string `json:"description"`
	}
)

func ToResponseCourses(courses []*entities.Course) []*Course {
	var results []*Course

	for _, course := range courses {
		results = append(results, &Course{
			ID:                 course.ID,
			CategoryId:         course.CategoryId,
			MentorId:           course.UserId,
			InterestId:         course.InterestId,
			Title:              course.Title,
			Category:           course.Category.Name,
			Interest:           course.Interest.Name,
			Mentor:             course.User.Fullname,
			Level:              course.Level,
			Type:               course.Type,
			IsPublished:        course.IsPublished,
			Price:              course.Price,
			Thumbnail:          course.Thumbnail,
			CategoryImage:      course.Category.Image,
			MentorProfileImage: course.User.ProfilePicture,
			CreatedAt:          utils.Timestamp(course.CreatedAt),
			UpdatedAt:          utils.Timestamp(course.UpdatedAt),
		})
	}

	return results
}

func ToResponseCourse(course *entities.Course) *Course {
	var modules []*Module

	for _, module := range course.Modules {
		modules = append(modules, &Module{
			ID:          module.ID,
			CourseId:    module.CourseId,
			Title:       module.Title,
			Description: module.Description,
			Materials: func(materials []*entities.Material) []*Material {
				var m []*Material

				for _, material := range materials {
					m = append(m, &Material{
						ID:          material.ID,
						CourseId:    material.CourseId,
						ModuleId:    material.ModuleId,
						Title:       material.Title,
						Url:         material.Url,
						Description: material.Description,
					})
				}

				return m
			}(module.Materials),
		})
	}

	return &Course{
		ID:                 course.ID,
		CategoryId:         course.CategoryId,
		MentorId:           course.UserId,
		InterestId:         course.InterestId,
		Title:              course.Title,
		Category:           course.Category.Name,
		Interest:           course.Interest.Name,
		Mentor:             course.User.Fullname,
		Level:              course.Level,
		Type:               course.Type,
		IsPublished:        course.IsPublished,
		Price:              course.Price,
		Thumbnail:          course.Thumbnail,
		CategoryImage:      course.Category.Image,
		MentorProfileImage: course.User.ProfilePicture,
		CreatedAt:          utils.Timestamp(course.CreatedAt),
		UpdatedAt:          utils.Timestamp(course.UpdatedAt),
		Modules:            modules,
	}
}

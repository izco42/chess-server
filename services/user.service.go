package services

import (
	"chess-server/database"
	"chess-server/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetAllUsers() []models.User {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println("Error al obtener usuarios:", result.Error)
	}
	log.Println("Usuarios obtenidos:", users)
	return users
}

func CreateUser(user models.User) (*models.User, error) {

	var existingUserByUsername models.User
	result := database.DB.Where("username = ?", user.Username).First(&existingUserByUsername)
	if result.Error == nil {
		log.Println("El username ya está en uso:", user.Username)
		return nil, errors.New("el username ya está en uso")
	}

	var existingUserByEmail models.User
	result = database.DB.Where("email = ?", user.Email).First(&existingUserByEmail)
	if result.Error == nil {
		log.Println("El email ya está en uso:", user.Email)
		return nil, errors.New("el email ya está en uso")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error al encriptar la contraseña:", err)
		return nil, err
	}
	user.Password = string(hashedPassword)

	result = database.DB.Create(&user)
	if result.Error != nil {
		log.Println("Error al crear el usuario:", result.Error)
		return nil, result.Error
	}

	log.Println("Usuario creado:", user)
	return &user, nil
}

func LoginUser(userLogin models.UserLogin) (*models.User, error) {
	user := new(models.User)

	result := database.DB.Where("email = ?", userLogin.Email).First(user)
	if result.Error != nil {
		log.Println("Error al encontrar el usuario:", result.Error)
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		log.Println("Contraseña incorrecta:", err)
		return nil, errors.New("contraseña incorrecta")
	}

	log.Println("Usuario autenticado con éxito:", user)
	return user, nil
}

func DeleteUser(id uint) error {

	var existingUser models.User
	result := database.DB.Where("id = ?", id).First(&existingUser)
	if result.Error != nil {
		log.Println("Error al encontrar el usuario:", result.Error)
		return result.Error
	}
	
	result = database.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		log.Println("Error al eliminar el usuario:", result.Error)
		return result.Error
	}
	log.Println("Usuario eliminado con éxito")
	return nil
}

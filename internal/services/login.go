package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/AidarIlyasov/whatsapp_app/internal/bootstrap"
	"github.com/Rhymen/go-whatsapp"
	"github.com/spf13/viper"
)

// login handles WhatsApp login or restores a saved session
func Login(app *bootstrap.App) error {
	// Attempt to restore the session from the CSV file
	_, err := restoreSession(app.Conn)
	if err == nil {
		fmt.Println("Session restored successfully!")
		return nil
	}

	// If no session found or restoration fails, perform a new login
	fmt.Println("No valid session found, logging in with QR code...")
	qr := make(chan string)
	go func() {
		for qrCode := range qr {
			fmt.Printf("Scan the QR Code to login: %s\n", qrCode)
		}
	}()

	session, err := app.Conn.Login(qr)
	if err != nil {
		return fmt.Errorf("error logging in: %w", err)
	}

	// Save the session to the CSV file
	if err := saveSession(session); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	fmt.Printf("Login successful! Session saved to file.\n")
	return nil
}

// saveSession saves the session to a CSV file
func saveSession(session whatsapp.Session) error {
	// Get the file path from Viper configs
	sessionFile := viper.GetString("sessions.path")
	if sessionFile == "" {
		return errors.New("session file path not configured")
	}

	file, err := os.Create(sessionFile)
	if err != nil {
		return fmt.Errorf("failed to create session file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write session details to the CSV file
	err = writer.Write([]string{
		session.ClientId,
		session.ServerToken,
		session.ClientToken,
		session.Wid,
	})

	if err != nil {
		return fmt.Errorf("failed to write session to file: %w", err)
	}

	return nil
}

// restoreSession restores the session from the CSV file
func restoreSession(wac *whatsapp.Conn) (whatsapp.Session, error) {
	// Get the file path from Viper configs
	sessionFile := viper.GetString("sessions.path")
	if sessionFile == "" {
		return whatsapp.Session{}, errors.New("session file path not configured")
	}

	file, err := os.Open(sessionFile)
	if err != nil {
		return whatsapp.Session{}, fmt.Errorf("failed to open session file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.Read()
	if err != nil {
		return whatsapp.Session{}, fmt.Errorf("failed to read session file: %w", err)
	}

	// Restore session from CSV file data
	session := whatsapp.Session{
		ClientId:    record[0],
		ServerToken: record[1],
		ClientToken: record[2],
		Wid:         record[3],
	}

	// Restore the session in WhatsApp connection
	_, err = wac.RestoreWithSession(session)
	if err != nil {
		return whatsapp.Session{}, fmt.Errorf("failed to restore session: %w", err)
	}

	return session, nil
}

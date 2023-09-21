package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Text RPG Game")

    // Create widgets for the game menu
    characterInfoBtn := widget.NewButton("Display Character Info", func() {
        // Implement the logic to display character info here
    })
    accessInventoryBtn := widget.NewButton("Access Inventory", func() {
        // Implement the logic to access inventory here
    })
    merchantBtn := widget.NewButton("Merchant", func() {
        // Implement the logic for the merchant menu here
    })
    blacksmithBtn := widget.NewButton("Blacksmith", func() {
        // Implement the logic for the blacksmith menu here
    })
    exitBtn := widget.NewButton("Exit", func() {
        myApp.Quit()
    })

    // Create a menu container
    menu := container.NewVBox(
        characterInfoBtn,
        accessInventoryBtn,
        merchantBtn,
        blacksmithBtn,
        exitBtn,
    )

    // Create a main game screen
    gameScreen := widget.NewLabel("Your game content goes here.")

    // Create a layout for the main window
    content := fyne.NewContainerWithLayout(
        layout.NewBorderLayout(nil, nil, menu, nil),
        menu,
        gameScreen,
    )

    // Set the main content of the window
    myWindow.SetContent(content)

    // Show the window
    myWindow.ShowAndRun()
}
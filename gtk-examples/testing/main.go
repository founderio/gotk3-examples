package main

import (
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	log.Println("Building Window")
	win := buildWindow()

	log.Println("Waiting for window to finish drawing")
	gtk.TestWidgetWaitForDraw(win)

	go func() {
		log.Println("Sleep 4 seconds")
		time.Sleep(time.Second * 4)

		glib.IdleAdd(func() {
			log.Println("Find button")
			widget, err := gtk.TestFindWidget(win, "Clicked Automatically", glib.TypeFromName("GtkButton"))
			if err != nil {
				panic(err)
			}

			button, ok := widget.(*gtk.Button)
			if !ok {
				panic("Not a button")
			}

			log.Println("Click button")
			button.Clicked()
			// We could do this, but I have not seen this work reliably
			//gtk.TestWidgetClick(button, 1, gdk.ModifierType(0))
		})

		log.Println("Sleep 2 seconds")
		time.Sleep(time.Second * 2)

		glib.IdleAdd(func() {
			win.Destroy()
			gtk.MainQuit()
		})
	}()

	gtk.Main()

	log.Println("Done")
}

func buildWindow() *gtk.Window {

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Testing Example")

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Waiting for something to happen...")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	box.PackStart(l, true, true, 0)

	button, err := gtk.ButtonNewWithLabel("Clicked Automatically")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	button.Connect("clicked", func() {
		l.SetText("Something did happen!")
	})
	box.PackStart(button, true, true, 0)

	// Add the label to the window.
	win.Add(box)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	return win
}

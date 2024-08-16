import tkinter as tk
from enum import Enum, auto


# [x]: Basic movement 
# [ ]: Font family and Font size
# [ ]: Be able to open text-file
# [ ]: Be able to save text-file 
# [ ]: Highlight keywords
# [ ]: Be able to keep track of stdout, stdin, and stderr
# [ ]: Be able to jump an error has been caused location of file.
# [ ]: Be able to put Todo area
# [ ]: Dark mode

class Direction(Enum):
    UP = auto()
    DOWN = auto()
    LEFT = auto()
    RIGHT = auto()

class TextField(tk.Text):
    def __init__(self, master, status_label):
        super().__init__(master, wrap="word", undo=True)
        self.status_label = status_label
        self.bind("<Control-Key-k>", self.on_pressed_movement_key(Direction.UP))
        self.bind("<Control-Key-j>", self.on_pressed_movement_key(Direction.DOWN))
        self.bind("<Control-Key-h>", self.on_pressed_movement_key(Direction.LEFT))
        self.bind("<Control-Key-l>", self.on_pressed_movement_key(Direction.RIGHT))

        self.bind("<Control_L>", self.on_ctrl_press)
        self.bind("<Control_R>", self.on_ctrl_press)
        self.bind("<KeyRelease-Control_L>", self.on_ctrl_release)
        self.bind("<KeyRelease-Control_R>", self.on_ctrl_release)
        self.direction_mapping = {
            Direction.UP:    lambda r, c: (max(r-1, 1), c),
            Direction.DOWN:  lambda r, c: (r+1, c),
            Direction.LEFT:  lambda r, c: (r, max(c-1, 0)),
            Direction.RIGHT: lambda r, c: (r, c+1),
        }

        self.press_ctrl = False

    def on_pressed_movement_key(self, direction):
        def inner(event):
            self.update_cursor(direction)
            return "break"
        return inner

    def on_ctrl_press(self, event):
        if not self.press_ctrl:
            self.status_label.config(text="---  MOVE  MODE ---")
            self.press_ctrl = True

    def on_ctrl_release(self, event):
        if self.press_ctrl:
            self.status_label.config(text="--- INSERT MODE ---")
            self.press_ctrl = False

    def update_cursor(self, direction):
        row, col = list(map(int, self.index(tk.INSERT).split('.')))
        new_row, new_col = self.direction_mapping[direction](row, col)
        self.mark_set(tk.INSERT, f"{new_row}.{new_col}")


root = tk.Tk()
root.title = "Text Editor"
root.geometry("1280x720")
status_label = tk.Label(root, text="--- INSERT MODE ---", anchor='w', relief=tk.SUNKEN)
status_label.pack(side=tk.BOTTOM, fill=tk.X)
text_field = TextField(root, status_label)
text_field.pack(expand="yes", fill="both")
text_field.focus()
root.mainloop()


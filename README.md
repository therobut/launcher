# Launcher
An extremely simple program that simply launches another program.

## Motivation
The Windows 10 start menu does not allow you to have multiple shortcuts to the same executable, even if they have different arguments. I have two Firefox profiles that I use regularly, and it takes (in my opinion) too many steps to launch a second profile.

With this program, I can compile as many copies as I want and simply make sure they all have a different executable name so they all can show up in the Windows 10 start menu.

## Directions
This program is not intended to be installed like a normal go program. The intention is that you compile a new executable for every shortcut you want to add to the start menu.

### Basic Example
To make a shortcut for a second Firefox profile (called "work-profile" in this example)

1. Download the launcher from the [releases page](https://github.com/therobut/launcher/releases)
2. Rename the exectuable to a unique name. For our example: `workFirefoxProfile.exe`
3. Right-click the executable and click `Send to > Desktop (create shortcut)`
4. Navigate to your desktop
5. Right-click the new shortcut and click `Properties`
6. Change the `Target` field to something like this: `"path-to-workFirefoxProfile.exe" "C:\Program Files\Mozilla Firefox\firefox.exe" -P "work-profile"`
7. Press the keys `WIN + R` and type `%appdata%\Microsoft\Windows\Start Menu\Programs` then press enter
8. Copy your new shortcut into this folder

Now your new shortcut should show up in your start menu!

### Build with icon
To build the executable with an icon, do the following

1. Outside of the launcher source directory, run `go get github.com/akavel/rsrc`
2. navigate back to the launcher source directory
3. copy your *.ico file into the launcher source directory
4. run `rsrc -ico ICON_FILE.ico`
5. build the executable `go build -o launcher.exe`

rsrc generates a `.syso` file which includes information about the icon. Go will automatically use the file when building the executable.
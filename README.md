# hypr-kbd-layout

Plugin for waybar to display the current keyboard layout.

## Installation
Prebuild binaries can be found at https://github.com/labi-le/hyrp-kbd-layout/releases/
```bash
sudo cp hypr-kbd-layout /usr/bin/
```

#### Compile manually

```bash
sudo make install
```

## Configuration

add the following to your waybar config:

```json
"custom/keyboard-layout": {
  "return-type": "json",
  "exec": "hypr-kbd-layout", 
  "format": "{}"
},
```
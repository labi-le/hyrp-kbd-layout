# hypr-kbd-layout

Plugin for waybar to display the current keyboard layout.

## Installation

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
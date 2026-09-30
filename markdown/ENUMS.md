# Enum Types in Kwenda

Kwenda supports enum (enumeration) types using the `aina` keyword (Swahili for "type" or "kind").

## Syntax

```swahili
aina EnumName {
    VALUE1,
    VALUE2,
    VALUE3
}
```

## Features

### Defining Enums

Enums are defined using the `aina` keyword followed by the enum name and a list of comma-separated values enclosed in braces:

```swahili
aina Status {
    PENDING,
    ACTIVE,
    COMPLETED,
    CANCELLED
}

aina Color {
    RED,
    GREEN,
    BLUE,
    YELLOW
}

aina Priority {
    LOW,
    MEDIUM,
    HIGH,
    URGENT
}
```

### Accessing Enum Values

Enum values are accessed using dot notation:

```swahili
current_status = Status.ACTIVE
my_color = Color.BLUE
task_priority = Priority.HIGH
```

### Printing Enum Values

Enum values are displayed in the format `EnumName.VALUE`:

```swahili
andika("Status: ", current_status)
# Output: Status: Status.ACTIVE

andika("Color: ", my_color)
# Output: Color: Color.BLUE
```

### Comparing Enum Values

Enums support equality (`==`) and inequality (`!=`) comparisons:

```swahili
aina Status {
    PENDING,
    ACTIVE,
    COMPLETED
}

current = Status.ACTIVE

# Equality comparison
kama current == Status.ACTIVE {
    andika("Status is ACTIVE")
}

# Inequality comparison
kama current != Status.PENDING {
    andika("Status is not PENDING")
}
```

### Using Enums in Conditionals

Enums are commonly used in conditional statements to represent discrete states:

```swahili
aina TrafficLight {
    RED,
    YELLOW,
    GREEN
}

light = TrafficLight.RED

kama light == TrafficLight.RED {
    andika("Stop!")
}

kama light == TrafficLight.YELLOW {
    andika("Caution!")
}

kama light == TrafficLight.GREEN {
    andika("Go!")
}
```

## Complete Example

```swahili
# Define day of week enum
aina Day {
    MONDAY,
    TUESDAY,
    WEDNESDAY,
    THURSDAY,
    FRIDAY,
    SATURDAY,
    SUNDAY
}

# Use enum values
today = Day.WEDNESDAY
tomorrow = Day.THURSDAY

andika("Today is: ", today)
andika("Tomorrow is: ", tomorrow)

# Compare enum values
kama today == Day.WEDNESDAY {
    andika("It's the middle of the week!")
}

kama tomorrow == Day.FRIDAY {
    andika("Almost the weekend!")
}
```

## Common Use Cases

### 1. Status Management
```swahili
aina OrderStatus {
    PENDING,
    PROCESSING,
    SHIPPED,
    DELIVERED,
    CANCELLED
}

order_status = OrderStatus.PROCESSING
```

### 2. Configuration Options
```swahili
aina LogLevel {
    DEBUG,
    INFO,
    WARNING,
    ERROR
}

current_level = LogLevel.INFO
```

### 3. Game States
```swahili
aina GameState {
    MENU,
    PLAYING,
    PAUSED,
    GAME_OVER
}

state = GameState.PLAYING
```

### 4. Directions
```swahili
aina Direction {
    NORTH,
    SOUTH,
    EAST,
    WEST
}

player_direction = Direction.NORTH
```

## Best Practices

1. **Use UPPERCASE for enum values** - This is the conventional style that makes enum values easily distinguishable from variables.

2. **Choose descriptive enum names** - The enum name should clearly describe what the values represent.

3. **Group related constants** - Use enums to group related constant values together rather than defining them as separate variables.

4. **Use enums for finite sets** - Enums are perfect for representing a fixed set of possible values (like days of the week, states, etc.).

## Notes

- Enum values are compared by both their enum type and value name
- Two enum values are equal only if they have the same enum type AND the same value name
- Enum values cannot be used in arithmetic operations
- Each enum type is independent - `Status.ACTIVE` is different from `Mode.ACTIVE` even if both have an `ACTIVE` value

## Implementation Details

- **Keyword**: `aina` (Swahili for "type/kind")
- **Access Pattern**: `EnumName.VALUE`
- **Supported Operations**: `==` (equality), `!=` (inequality)
- **Display Format**: `EnumName.VALUE`

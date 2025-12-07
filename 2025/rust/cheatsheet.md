# Rust Cheat Sheet für Advent of Code

## Grundstruktur einer Lösung

```rust
use std::collections::{HashMap, HashSet, VecDeque};

fn solve(input: &str) -> i64 {
    let mut result = 0;
    
    for line in input.lines() {
        // ...
    }
    
    result
}
```

---

## Variablen & Typen

```rust
let x = 42;              // immutable, Typ wird inferiert
let mut y = 0;           // mutable
let z: i64 = 100;        // expliziter Typ

let s = String::from("hello");   // owned String
let s: &str = "hello";           // string slice (borrowed)
```

**Go → Rust Typen:**
| Go | Rust |
|-----|------|
| `int` | `i32` oder `i64` |
| `string` | `String` (owned) oder `&str` (borrowed) |
| `[]int` | `Vec<i32>` |
| `map[string]int` | `HashMap<String, i32>` |
| `bool` | `bool` |

---

## Strings & Parsing

### String zu Zahl (Go: `strconv.Atoi`)
```rust
let n: i32 = "42".parse().unwrap();
let n = "42".parse::<i64>().unwrap();

// Fehlerbehandlung
if let Ok(n) = "42".parse::<i32>() {
    println!("{}", n);
}
```

### String splitten (Go: `strings.Split`)
```rust
let parts: Vec<&str> = "a,b,c".split(',').collect();
let parts: Vec<&str> = "a  b  c".split_whitespace().collect();

// Direktes Iterieren (ohne collect)
for part in "a,b,c".split(',') {
    println!("{}", part);
}

// In genau 2 Teile splitten
let (left, right) = "key=value".split_once('=').unwrap();
```

### String-Prüfungen (Go: `strings.HasPrefix`, `Contains`)
```rust
s.starts_with("foo")    // HasPrefix
s.ends_with("bar")      // HasSuffix
s.contains("baz")       // Contains
s.trim()                // TrimSpace
s.replace("a", "b")     // ReplaceAll
```

### Zeilen iterieren (Go: `strings.Split(s, "\n")`)
```rust
for line in input.lines() {
    // line ist &str
}
```

### Chars iterieren
```rust
for c in s.chars() {
    // c ist char
}

// Mit Index
for (i, c) in s.chars().enumerate() {
    println!("{}: {}", i, c);
}

// Char zu Digit
let digit = c.to_digit(10).unwrap();  // -> u32
```

---

## Vektoren (Go: Slices)

```rust
let mut v: Vec<i32> = Vec::new();
let mut v = vec![1, 2, 3];           // Mit Werten initialisieren

v.push(4);                           // append
v.pop();                             // Letztes Element entfernen
v.len();                             // Länge
v.is_empty();                        // len == 0

v[0]                                 // Zugriff (panics wenn out of bounds)
v.get(0)                             // -> Option<&T>

v.sort();                            // In-place sortieren
v.reverse();                         // In-place umkehren
v.contains(&42);                     // Enthält Element?

// Iterieren
for x in &v {         // &v = Referenz, v bleibt nutzbar
    println!("{}", x);
}

for x in &mut v {     // Mutable Referenz
    *x += 1;
}
```

---

## HashMap (Go: `map`)

```rust
use std::collections::HashMap;

let mut map: HashMap<String, i32> = HashMap::new();

map.insert("foo".to_string(), 42);
map.get("foo")                      // -> Option<&i32>
map.get("foo").unwrap_or(&0)        // Mit Default
map.contains_key("foo")             // -> bool
map.remove("foo");

// Entry API (sehr nützlich für Zähler)
*map.entry("foo".to_string()).or_insert(0) += 1;

// Iterieren
for (key, value) in &map {
    println!("{}: {}", key, value);
}
```

---

## HashSet

```rust
use std::collections::HashSet;

let mut set: HashSet<i32> = HashSet::new();

set.insert(42);                     // -> bool (true wenn neu)
set.contains(&42);                  // -> bool
set.remove(&42);
set.len();
```

---

## Option & Result

### Option (Go: Rückgabe von `value, ok`)
```rust
let maybe: Option<i32> = Some(42);
let nothing: Option<i32> = None;

// Unwrap (panics wenn None)
let x = maybe.unwrap();

// Mit Default
let x = maybe.unwrap_or(0);
let x = maybe.unwrap_or_default();  // 0 für Zahlen, "" für String

// Pattern Matching
if let Some(x) = maybe {
    println!("{}", x);
}

match maybe {
    Some(x) => println!("{}", x),
    None => println!("nichts"),
}
```

### Result (Go: `value, err`)
```rust
let result: Result<i32, String> = Ok(42);
let error: Result<i32, String> = Err("failed".to_string());

// Unwrap (panics bei Err)
let x = result.unwrap();

// Fehler weiterreichen (nur in Funktionen die Result zurückgeben)
let x = result?;

// Pattern Matching
match result {
    Ok(x) => println!("{}", x),
    Err(e) => println!("Error: {}", e),
}
```

---

## Match (Go: `switch`)

```rust
match value {
    0 => println!("null"),
    1 | 2 => println!("eins oder zwei"),
    3..=9 => println!("drei bis neun"),
    n if n < 0 => println!("negativ: {}", n),
    _ => println!("was anderes"),
}

// Match mit Rückgabewert
let result = match value {
    0 => "null",
    _ => "nicht null",
};

// Match auf Tuple
match (x, y) {
    (0, 0) => println!("Ursprung"),
    (0, _) => println!("Y-Achse"),
    (_, 0) => println!("X-Achse"),
    _ => println!("irgendwo"),
}
```

---

## If Let / While Let

```rust
// Statt match für einzelnen Fall
if let Some(x) = maybe_value {
    println!("{}", x);
}

// Kombiniert mit else
if let Some(x) = maybe_value {
    println!("{}", x);
} else {
    println!("war None");
}

// While let - iteriert bis None
while let Some(x) = stack.pop() {
    println!("{}", x);
}
```

---

## Loops

```rust
// For mit Range (Go: for i := 0; i < 10; i++)
for i in 0..10 {        // 0 bis 9
    println!("{}", i);
}

for i in 0..=10 {       // 0 bis 10 (inklusiv)
    println!("{}", i);
}

// Rückwärts
for i in (0..10).rev() {
    println!("{}", i);
}

// Endlosschleife (Go: for {})
loop {
    if condition {
        break;
    }
}

// Loop mit Rückgabewert
let result = loop {
    if found {
        break 42;
    }
};

// While
while condition {
    // ...
}
```

---

## Iteratoren (funktionaler Stil)

```rust
// Map (transformieren)
let doubled: Vec<i32> = v.iter().map(|x| x * 2).collect();

// Filter
let evens: Vec<i32> = v.iter().filter(|x| *x % 2 == 0).copied().collect();

// Filter + Map kombiniert
let parsed: Vec<i32> = lines
    .iter()
    .filter_map(|s| s.parse().ok())
    .collect();

// Enumerate (Index + Wert)
for (i, x) in v.iter().enumerate() {
    println!("{}: {}", i, x);
}

// Sum
let total: i32 = v.iter().sum();

// Count
let count = v.iter().filter(|x| **x > 0).count();

// Find (erstes Element das Bedingung erfüllt)
let found = v.iter().find(|x| **x > 10);  // -> Option<&i32>

// Any / All
let has_positive = v.iter().any(|x| *x > 0);
let all_positive = v.iter().all(|x| *x > 0);

// Min / Max
let min = v.iter().min();   // -> Option<&i32>
let max = v.iter().max();
```

---

## Funktionen

```rust
fn add(a: i32, b: i32) -> i32 {
    a + b    // Kein Semicolon = Return
}

fn add(a: i32, b: i32) -> i32 {
    return a + b;    // Explizites return auch möglich
}

// Mehrere Rückgabewerte (Go-Style) -> Tuple
fn parse_pair(s: &str) -> (i32, i32) {
    let parts: Vec<&str> = s.split(',').collect();
    (parts[0].parse().unwrap(), parts[1].parse().unwrap())
}

let (x, y) = parse_pair("10,20");

// Option als Rückgabe (statt Go's bool)
fn find_value(s: &str) -> Option<i32> {
    if s.is_empty() {
        None
    } else {
        Some(s.parse().unwrap())
    }
}
```

---

## File I/O

```rust
use std::fs;

// Ganze Datei lesen (Go: os.ReadFile)
let content = fs::read_to_string("input.txt").unwrap();

// In deinem AoC-Setup:
let input = read_input(1);  // Aus lib.rs
```

---

## Regex

```rust
use regex::Regex;

let re = Regex::new(r"(\d+)").unwrap();

// Finden ob Match existiert
if re.is_match(text) { ... }

// Erstes Match finden
if let Some(caps) = re.captures(text) {
    let num: i32 = caps[1].parse().unwrap();
}

// Alle Matches finden
for caps in re.captures_iter(text) {
    let num: i32 = caps[1].parse().unwrap();
}

// Alle Matches als Iterator
let numbers: Vec<i32> = re
    .captures_iter(text)
    .map(|c| c[1].parse().unwrap())
    .collect();
```

---

## 2D Grid (häufig bei AoC)

```rust
// Grid als Vec<Vec<char>>
let grid: Vec<Vec<char>> = input
    .lines()
    .map(|line| line.chars().collect())
    .collect();

let height = grid.len();
let width = grid[0].len();

// Zugriff
let c = grid[y][x];

// Bounds Check
if x >= 0 && x < width as i32 && y >= 0 && y < height as i32 {
    let c = grid[y as usize][x as usize];
}

// Nachbarn (4 Richtungen)
let dirs = [(0, -1), (1, 0), (0, 1), (-1, 0)];  // N, E, S, W

for (dx, dy) in dirs {
    let nx = x + dx;
    let ny = y + dy;
    if let Some(row) = grid.get(ny as usize) {
        if let Some(&c) = row.get(nx as usize) {
            // c ist der Nachbar
        }
    }
}

// Oder nutze die Grid-Struct aus lib.rs!
let grid = Grid::from_string(&input);
if let Some(&c) = grid.get(x, y) {
    // ...
}
```

---

## BFS / Queue

```rust
use std::collections::VecDeque;

let mut queue = VecDeque::new();
queue.push_back(start);

while let Some(current) = queue.pop_front() {
    // Verarbeite current
    for next in neighbors(current) {
        queue.push_back(next);
    }
}
```

---

## Visited Set Pattern

```rust
use std::collections::HashSet;

let mut visited = HashSet::new();
let mut queue = VecDeque::new();

queue.push_back(start);
visited.insert(start);

while let Some(current) = queue.pop_front() {
    for next in neighbors(current) {
        if visited.insert(next) {  // insert gibt false zurück wenn schon drin
            queue.push_back(next);
        }
    }
}
```

---

## Häufige AoC Patterns

### Input parsen: Zahlen aus jeder Zeile
```rust
let numbers: Vec<i64> = input
    .lines()
    .map(|line| line.parse().unwrap())
    .collect();
```

### Input parsen: Mehrere Zahlen pro Zeile
```rust
for line in input.lines() {
    let nums: Vec<i32> = line
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();
}
```

### Input parsen: Key-Value Format
```rust
for line in input.lines() {
    let (key, value) = line.split_once(": ").unwrap();
    let value: i32 = value.parse().unwrap();
}
```

### Input parsen: Durch Leerzeilen getrennte Blöcke
```rust
let blocks: Vec<&str> = input.split("\n\n").collect();

for block in input.split("\n\n") {
    for line in block.lines() {
        // ...
    }
}
```

### Zählen
```rust
let count = input.lines().filter(|line| is_valid(line)).count();
```

### Maximum finden
```rust
let max = input
    .lines()
    .map(|line| calculate(line))
    .max()
    .unwrap();
```

---

## Schnelle Referenz: Die wichtigsten 20 Dinge

| Was | Rust |
|-----|------|
| String zu i32 | `s.parse::<i32>().unwrap()` |
| i32 zu String | `n.to_string()` |
| String splitten | `s.split(',')` oder `s.split_whitespace()` |
| Zeilen iterieren | `input.lines()` |
| Chars iterieren | `s.chars()` |
| Vec erstellen | `Vec::new()` oder `vec![1,2,3]` |
| Vec anhängen | `v.push(x)` |
| HashMap erstellen | `HashMap::new()` |
| HashMap einfügen | `map.insert(k, v)` |
| HashMap lesen | `map.get(&k)` |
| HashMap Zähler | `*map.entry(k).or_insert(0) += 1` |
| HashSet | `set.insert(x)`, `set.contains(&x)` |
| Option auspacken | `.unwrap()` oder `if let Some(x) = ...` |
| For-Schleife | `for i in 0..n { }` |
| Iterator zu Vec | `.collect::<Vec<_>>()` oder `.collect()` |
| Transformieren | `.map(\|x\| ...)` |
| Filtern | `.filter(\|x\| ...)` |
| Summieren | `.sum::<i32>()` |
| Sortieren | `v.sort()` |
| Datei lesen | `fs::read_to_string("file.txt")` |

---

## Typische Fehler & Fixes

**"borrowed value does not live long enough"**
→ Sammle in einen owned Vec: `.map(|s| s.to_string()).collect()`

**"cannot move out of borrowed content"**
→ Nutze `.clone()` oder arbeite mit Referenzen

**"expected `&str`, found `String`"**
→ Nutze `&s` oder `s.as_str()`

**"expected `String`, found `&str`"**
→ Nutze `s.to_string()` oder `String::from(s)`

**Iterator liefert Referenzen statt Werte**
→ Füge `.copied()` oder `.cloned()` hinzu

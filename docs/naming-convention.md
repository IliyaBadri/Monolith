# General Naming Conventions (GO)

## 1. CamelCase for identifiers
- Use **mixedCaps** or **camelCase** (not underscores).
- Start with an **uppercase** letter for **exported** names (public).
- Start with a **lowercase** letter for **unexported** names (private to the package).

```go
// Exported (visible outside the package)
func PrintMessage() {}

// Unexported (private to the package)
func printMessage() {}
```

## 2.  **Package Naming**
- Use **lowercase**, **single-word**, **no underscores**, **no plurals**.
- Name the package after its functionality.

```go
package math     // Good
package strings  // Good

package utils    // Avoid generic names
package my_utils // Avoid underscores
```

## 3. **File Naming**
- All lowercase, words separated by **underscores** are okay for file names.
```go
math_utils.go     // Acceptable
http_client.go    // Acceptable
```

## 4. Constants
- Use **CamelCase**, often all uppercase is **not idiomatic** in Go.

```go
const Pi = 3.14       // Idiomatic
const MAX_VALUE = 10  // Not idiomatic
```

## 5. Interfaces
- If an interface has a single method, name it by the method + “er”
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Formatter interface {
    Format() string
}
```

## 6. Acronyms
- Use **consistent casing**:

```go
type HTTPClient struct {}  // Good
type HttpClient struct {}  // Less idiomatic
```
# General Naming Conventions (HTML)

## 1. Element IDs and class names:
 - Use **kebab-case**:
```html
<div class="element-id-1"></div> <-- Good -->
<input id="username-input"/>
```

# General Naming Conventions (JS)

## 1. Variables and Functions:
 - Use **camelCase**

```js
let customEventManager = Object(); // Good
```

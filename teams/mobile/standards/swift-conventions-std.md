# Swift Conventions Standard

## Scope
This standard applies to all Swift code in iOS applications.

## Rules

### 1. Naming Conventions

**Rule:** Use camelCase for variables and functions, PascalCase for types.

**Rationale:** Follows Swift conventions and improves readability.

**Example:**
```swift
// Good
var userName: String
func fetchUserData() -> User
class UserProfileViewController: UIViewController

// Bad
var user_name: String
func FetchUserData() -> User
class userProfileViewController: UIViewController
```

### 2. Optionals

**Rule:** Avoid force unwrapping. Use optional binding or guard statements.

**Rationale:** Force unwrapping causes crashes at runtime.

**Example:**
```swift
// Good
if let name = userName {
    print(name)
}

guard let data = responseData else {
    return
}

// Bad
let name = userName!
```

### 3. Access Control

**Rule:** Use private by default, open only when necessary.

**Rationale:** Minimizes API surface and prevents misuse.

**Example:**
```swift
// Good
private func processData() -> String

// Only when subclassing outside module
open class BaseViewController: UIViewController
```

### 4. Closure Memory Management

**Rule:** Use capture lists to prevent retain cycles.

**Rationale:** Prevents memory leaks in closures.

**Example:**
```swift
// Good
viewModel.onComplete = { [weak self] result in
    guard let self = self else { return }
    self.updateUI(result)
}

// Bad
viewModel.onComplete = { result in
    self.updateUI(result)
}
```

### 5. SwiftUI View Composition

**Rule:** Extract complex views into separate components.

**Rationale:** Improves readability and reusability.

**Example:**
```swift
// Good
struct ProfileView: View {
    var body: some View {
        VStack {
            ProfileHeader(user: user)
            ProfileStats(stats: stats)
            ProfileActions(actions: actions)
        }
    }
}

// Bad
struct ProfileView: View {
    var body: some View {
        VStack {
            // 100 lines of UI code
        }
    }
}
```

### 6. Constants

**Rule:** Use struct for organizing constants.

**Rationale:** Groups related constants and provides namespace.

**Example:**
```swift
// Good
struct Constants {
    struct API {
        static let baseURL = "https://api.example.com"
        static let timeout: TimeInterval = 30
    }
}

// Usage
let url = Constants.API.baseURL

// Bad
let API_BASE_URL = "https://api.example.com"
```

### 7. Error Handling

**Rule:** Use Swift's error handling with do-catch.

**Rationale:** Proper error handling prevents crashes and aids debugging.

**Example:**
```swift
// Good
do {
    let data = try JSONEncoder().encode(user)
    try saveData(data)
} catch {
    print("Failed to save user: \(error)")
}

// Bad
let data = try! JSONEncoder().encode(user)
```

### 8. Extensions

**Rule:** Use extensions to organize code by functionality.

**Rationale:** Improves code organization and readability.

**Example:**
```swift
// Good
class UserViewController: UIViewController {
    // Main view controller code
}

// MARK: - UITableViewDataSource
extension UserViewController: UITableViewDataSource {
    // Table view data source methods
}

// MARK: - UITableViewDelegate
extension UserViewController: UITableViewDelegate {
    // Table view delegate methods
}
```

### 9. Async/Await

**Rule:** Use async/await for new code, prefer it over completion handlers.

**Rationale:** Modern concurrency is more readable and maintainable.

**Example:**
```swift
// Good (async/await)
func fetchUser() async throws -> User {
    let (data, _) = try await URLSession.shared.data(from: url)
    return try JSONDecoder().decode(User.self, from: data)
}

// Avoid (completion handler)
func fetchUser(completion: @escaping (Result<User, Error>) -> Void)
```

### 10. Comments

**Rule:** Document public APIs with documentation comments.

**Rationale:** Helps users of your code understand its usage.

**Example:**
```swift
/// Fetches user data from the server.
/// - Parameter userId: The unique identifier of the user
/// - Returns: A User object containing the user's data
/// - Throws: An error if the network request fails or data is invalid
func fetchUser(userId: String) async throws -> User {
    // Implementation
}
```

## Enforcement

### Linting
- SwiftLint configuration: `.swiftlint.yml`
- Run: `swiftlint`

### Review
All pull requests must:
- Pass SwiftLint checks
- Pass SwiftLint strict rules
- Have unit tests with >70% coverage
- Be reviewed by at least one team member

### Pre-commit Hooks
- SwiftLint runs automatically
- SwiftFormat runs automatically

## Exceptions

Exceptions require:
- Team lead approval
- Documentation of rationale
- Code comment explaining deviation

## Related Standards
- SwiftUI Conventions
- UIKit Conventions
- Testing Standards

## Tags
`mobile`, `ios`, `swift`, `standards`, `conventions`

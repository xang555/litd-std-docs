# Mobile Team Best Practices

## Overview
Best practices for mobile development (iOS and Android) at our company.

## Architecture

### MVVM Pattern

Use MVVM for consistent architecture:

```
View ←→ ViewModel ←→ Model
                  ↓
              Repository
```

### Dependency Injection

Use DI frameworks (SwiftDI on iOS, Hilt on Android):

```swift
// iOS
protocol UserRepository {
    func fetchUser(id: String) async throws -> User
}

final class UserProfileViewModel {
    private let repository: UserRepository

    init(repository: UserRepository) {
        self.repository = repository
    }
}
```

## iOS Best Practices

### SwiftUI

1. **Prefer SwiftUI for new features**
2. **Keep views small and focused**
3. **Use proper state management**
4. **Leverage previews for development**

```swift
struct UserRow: View {
    let user: User

    var body: some View {
        HStack {
            AsyncImage(url: user.avatarURL)
            VStack(alignment: .leading) {
                Text(user.name)
                Text(user.email).font(.caption)
            }
        }
    }
}
```

### Combine Framework

1. **Use Combine for reactive programming**
2. **Always cancel subscriptions in deinit**
3. **Use @Published for observable properties**

```swift
final class UserViewModel: ObservableObject {
    @Published var users: [User] = []
    private var cancellables = Set<AnyCancellable>()

    func loadUsers() {
        userRepository.fetchUsers()
            .receive(on: DispatchQueue.main)
            .sink { completion in
                // Handle completion
            } receiveValue: { [weak self] users in
                self?.users = users
            }
            .store(in: &cancellables)
    }
}
```

### Performance

1. **LazyVStack/LazyHStack for long lists**
2. **AsyncImage for remote images**
3. **Avoid doing work in body**

## Android Best Practices

### Jetpack Compose

1. **Prefer Compose for new UI**
2. **Keep composables small and reusable**
3. **Remember vs rememberSaveable**
4. **Side effects in LaunchedEffect**

```kotlin
@Composable
fun UserRow(user: User) {
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .padding(16.dp)
    ) {
        AsyncImage(
            model = user.avatarURL,
            contentDescription = null
        )
        Column {
            Text(text = user.name)
            Text(
                text = user.email,
                style = MaterialTheme.typography.caption
            )
        }
    }
}
```

### Coroutines

1. **Use coroutines for async operations**
2. **Structured concurrency**
3. **Proper scope selection**

```kotlin
class UserViewModel(
    private val repository: UserRepository
) : ViewModel() {
    private val _users = MutableStateFlow<List<User>>(emptyList())
    val users: StateFlow<List<User>> = _users.asStateFlow()

    fun loadUsers() {
        viewModelScope.launch {
            repository.fetchUsers()
                .collect { users ->
                    _users.value = users
                }
        }
    }
}
```

### Performance

1. **LazyColumn for long lists**
2. **Coil for image loading**
3. **Avoid recomposition**

## Cross-Platform Concerns

### Localization

Externalize all user-facing strings:

```swift
// iOS
Text("welcome_message")
// Localizable.strings
"welcome_message" = "Welcome!";
```

```kotlin
// Android
Text(stringResource(R.string.welcome_message))
```

### Accessibility

1. **Add content descriptions**
2. **Support dynamic type**
3. **Test with Voice Over/TalkBack**

```swift
// iOS
Image(systemName: "star")
    .accessibilityLabel("Favorite")
    .accessibilityHint("Marks as favorite")
```

```kotlin
// Android
Icon(
    imageVector = Icons.Default.Star,
    contentDescription = stringResource(R.string.favorite)
)
```

### Testing

1. **Unit tests for ViewModels**
2. **UI tests for critical flows**
3. **Screenshot tests for UI**

```swift
// iOS
func testUserViewModelLoadUsers() async throws {
    let mockRepository = MockUserRepository()
    mockRepository.users = [testUser]
    let viewModel = UserViewModel(repository: mockRepository)

    await viewModel.loadUsers()

    XCTAssertEqual(viewModel.users.count, 1)
}
```

```kotlin
// Android
@Test
fun `loadUsers updates users`() = runTest {
    val mockRepository = MockUserRepository()
    val viewModel = UserViewModel(mockRepository)

    viewModel.loadUsers()

    assertEquals(1, viewModel.users.value.size)
}
```

## Security

### API Keys

Never store API keys in code:

```swift
// iOS - Use environment variables or keychain
let apiKey = ProcessInfo.processInfo.environment["API_KEY"]!
```

```kotlin
// Android - Use BuildConfig or local.properties
BuildConfig.API_KEY
```

### Data Storage

1. **Use Keychain/EncryptedSharedPreferences for sensitive data**
2. **Never store passwords**
3. **Use App Transport Security (iOS)**

## Performance Monitoring

1. **Track app launch time**
2. **Monitor frame rates**
3. **Track API response times**
4. **Monitor crash rates**

## Resources

### iOS
- [SwiftUI Documentation](https://developer.apple.com/documentation/swiftui)
- [Human Interface Guidelines](https://developer.apple.com/design/human-interface-guidelines/)
- [Swift by Sundell](https://www.swiftbysundell.com/)

### Android
- [Jetpack Compose Documentation](https://developer.android.com/jetpack/compose)
- [Material Design](https://material.io/design)
- [Android Developers](https://developer.android.com/)

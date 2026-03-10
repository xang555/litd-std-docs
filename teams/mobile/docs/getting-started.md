# Mobile Team - Getting Started

## Welcome to the Mobile Team

This guide helps you get started with mobile development at our company.

## Tech Stack

### iOS
- **Language**: Swift
- **Framework**: SwiftUI
- **Architecture**: MVVM

### Android
- **Language**: Kotlin
- **Framework**: Jetpack Compose
- **Architecture**: MVVM

## Setup

### iOS

1. Install dependencies:
   ```bash
   cd ios && pod install
   ```

2. Open workspace:
   ```bash
   open App.xcworkspace
   ```

### Android

1. Open project in Android Studio
2. Sync Gradle files

## Project Structure

```
ios/
├── App/
│   ├── Views/          # SwiftUI views
│   ├── ViewModels/     # ViewModels
│   ├── Models/         # Data models
│   └── Services/       # API services
android/
├── app/
│   ├── presentation/   # UI components
│   ├── domain/         # Business logic
│   └── data/           # Data layer
```

## Code Style

Follow the coding standards defined in `teams/mobile/standards/`.

## Resources

- Team-specific agents: `teams/mobile/agents/`
- Team workflows: `teams/mobile/workflows/`
- Shared resources: `shared/`

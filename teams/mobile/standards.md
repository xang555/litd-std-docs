# Mobile Team Standards

## Overview

This document outlines the coding standards, best practices, and conventions for the Mobile Engineering team.

## Technology Stack

### iOS
- **Language**: Swift 5.9+
- **UI Framework**: SwiftUI
- **Architecture**: MVVM
- **Dependency Management**: Swift Package Manager
- **Testing**: XCTest

### Android
- **Language**: Kotlin 1.9+
- **UI Framework**: Jetpack Compose
- **Architecture**: MVI + Clean Architecture
- **Dependency Management**: Gradle (Kotlin DSL)
- **Testing**: JUnit5 + MockK

## Code Standards

### File Naming

- iOS: `PascalCase.swift` (e.g., `UserProfileView.swift`)
- Android: `PascalCase.kt` (e.g., `UserProfileScreen.kt`)

### Architecture Principles

- Single responsibility
- Dependency injection
- Unidirectional data flow
- Reactive programming with Combine/Flow

## Best Practices

- Screens should be < 300 lines
- Use view models for business logic
- Implement proper error handling
- Test UI components with previews/composables
- Follow platform design guidelines

## Related Documents

- [iOS Guidelines](./ios-guidelines.md)
- [Android Guidelines](./android-guidelines.md)

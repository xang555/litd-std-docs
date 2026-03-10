# Screen Create Command

## Description
Command template for scaffolding new mobile screens with MVVM architecture.

## Usage
```
/screen-create <ScreenName> [options]
```

## Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| ScreenName | string | true | - | Name of the screen (PascalCase) |
| --platform | string | false | both | Target platform: ios, android, both |
| --with-viewmodel | boolean | false | true | Include ViewModel |
| --with-repository | boolean | false | false | Include repository layer |
| --navigation | string | false | push | Navigation type: push, modal, tab |

## Examples

### iOS screen
```bash
/screen-create Profile --platform=ios
```
Creates iOS SwiftUI screen with ViewModel

### Android screen
```bash
/screen-create Settings --platform=android --navigation=modal
```
Creates Android Compose screen with modal navigation

### Cross-platform
```bash
/screen-create Dashboard --platform=both
```
Creates both iOS and Android screens

## Generated Structure

### iOS
```
Screens/ScreenName/
├── ScreenNameView.swift       # SwiftUI view
├── ScreenNameViewModel.swift  # ViewModel
├── ScreenNameModel.swift      # Data models
└── ScreenNameRouter.swift     # Navigation
```

### Android
```
screens/screenname/
├── ScreenNameScreen.kt        # Composable
├── ScreenNameViewModel.kt     # ViewModel
├── ScreenNameModel.kt         # Data models
└── ScreenNameNavigation.kt    # Navigation
```

## iOS Template

```swift
import SwiftUI

struct ScreenNameView: View {
    @StateObject private var viewModel = ScreenNameViewModel()

    var body: some View {
        ScrollView {
            VStack(spacing: 16) {
                // Screen content
            }
            .padding()
        }
        .navigationTitle("Screen Name")
        .onAppear {
            viewModel.onAppear()
        }
    }
}

#Preview {
    NavigationView {
        ScreenNameView()
    }
}
```

## Android Template

```kotlin
@Composable
fun ScreenNameScreen(
    viewModel: ScreenNameViewModel = viewModel()
) {
    val uiState by viewModel.uiState.collectAsState()

    Scaffold(
        topBar = {
            TopAppBar(
                title = { Text("Screen Name") }
            )
        }
    ) { padding ->
        LazyColumn(
            modifier = Modifier
                .fillMaxSize()
                .padding(padding),
            content = {
                // Screen content
            }
        )
    }

    LaunchedEffect(Unit) {
        viewModel.onAppear()
    }
}
```

## Tags
`mobile`, `ios`, `android`, `screen`, `scaffold`, `cli`

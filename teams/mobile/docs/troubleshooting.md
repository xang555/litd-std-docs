# Mobile Team Troubleshooting Guide

## Common Issues and Solutions

## iOS Build Issues

### "Command PhaseScriptExecution failed"

**Symptoms:**
Build fails during script phase

**Solutions:**
1. Clean build folder: Cmd+Shift+K
2. Clear derived data
3. Check CocoaPods installation: `pod install`
4. Verify Xcode version compatibility

### "Cannot find type 'X' in scope"

**Symptoms:**
Type not recognized, build errors

**Solutions:**
1. Clean build folder
2. Check import statements
3. Verify module membership
4. Restart Xcode

### SwiftUI Preview Not Working

**Symptoms:**
Previews fail to load

**Solutions:**
1. Check for macOS version compatibility
2. Verify Xcode is up to date
3. Try opening Canvas manually: Editor > Canvas
4. Check for errors in the view

## Android Build Issues

### "Execution failed for task ':app:mergeDebugResources'"

**Symptoms:**
Resource merge conflicts

**Solutions:**
1. Clean project: Build > Clean Project
2. Invalidate caches: File > Invalidate Caches
3. Check for duplicate resource names
4. Sync Gradle files

### "Could not resolve dependencies"

**Symptoms:**
Gradle sync fails

**Solutions:**
1. Check internet connection
2. Verify Maven repository URLs
3. Clear Gradle cache: `./gradlew clean build --refresh-dependencies`
4. Check proxy settings

### Compose Preview Not Working

**Symptoms:**
Previews fail to render

**Solutions:**
1. Enable Compose preview in settings
2. Check for compilation errors
3. Restart Android Studio
4. Update Compose version

## Runtime Issues

### App Crashes on Launch

**Symptoms:**
App crashes immediately after opening

**Solutions:**
1. Check crash logs (Xcode/Logcat)
2. Verify Info.plist/AndroidManifest.xml
3. Check for missing required permissions
4. Test with clean build

### Memory Leaks

**Symptoms:**
App slows down over time, crashes eventually

**Diagnosis:**
- iOS: Instruments > Leaks
- Android: Android Profiler > Memory

**Solutions:**
1. Fix retain cycles (use [weak self])
2. Unregister observers in deinit/onDestroy
3. Cancel network requests in deinit/onDestroy
4. Release resources when not needed

```swift
// iOS - Fix retain cycle
class UserProfileViewController: UIViewController {
    private var cancellables = Set<AnyCancellable>()

    override func viewDidLoad() {
        super.viewDidLoad()
        viewModel.$user
            .sink { [weak self] user in
                self?.updateUI(user)
            }
            .store(in: &cancellables)
    }
}
```

```kotlin
// Android - Fix memory leak
class UserProfileFragment : Fragment() {
    private var job: Job? = null

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        job = viewLifecycleOwner.lifecycleScope.launch {
            viewModel.user.collect { user ->
                updateUI(user)
            }
        }
    }

    override fun onDestroyView() {
        job?.cancel()
        super.onDestroyView()
    }
}
```

## UI Issues

### Layout Problems

**Symptoms:**
UI doesn't display correctly

**Solutions:**
1. Check constraints (iOS) or layouts (Android)
2. Test on different screen sizes
3. Verify safe area usage
4. Check for conflicting modifiers

### Animation Jank

**Symptoms:**
Animations are not smooth

**Solutions:**
1. Move heavy work off main thread
2. Use Instruments/Profiler to identify bottlenecks
3. Reduce animation complexity
4. Use proper animation APIs

## Network Issues

### API Calls Failing

**Symptoms:**
Network requests fail or timeout

**Solutions:**
1. Check network connectivity
2. Verify API endpoint URLs
3. Check authentication tokens
4. Enable network logging

```swift
// iOS - Debug network
let configuration = URLSessionConfiguration.default
configuration.httpAdditionalHeaders = ["Authorization": "Bearer \(token)"]
let session = URLSession(configuration: configuration)
```

```kotlin
// Android - Debug network
val okHttpClient = OkHttpClient.Builder()
    .addInterceptor { chain ->
        val request = chain.request().newBuilder()
            .addHeader("Authorization", "Bearer $token")
            .build()
        chain.proceed(request)
    }
    .build()
```

## Performance Issues

### Slow App Launch

**Diagnosis:**
- Measure launch time
- Profile main thread work

**Solutions:**
1. Defer non-critical work
2. Optimize app initialization
3. Reduce framework usage
4. Preload critical resources

### High Battery Usage

**Solutions:**
1. Optimize location updates
2. Use background tasks efficiently
3. Reduce unnecessary polling
4. Implement proper lifecycle handling

## Testing Issues

### Flaky Tests

**Symptoms:**
Tests pass/fail inconsistently

**Solutions:**
1. Add explicit waits/timing
2. Remove dependencies on timing
3. Mock external dependencies
4. Ensure test isolation

### UI Tests Failing

**Symptoms:**
UI tests can't find elements

**Solutions:**
1. Add accessibility identifiers
2. Wait for elements to appear
3. Check for timing issues
4. Verify test environment

## Deployment Issues

### App Store Rejection (iOS)

**Symptoms:**
App rejected by Apple

**Solutions:**
1. Review App Store Review Guidelines
2. Check for missing metadata
3. Verify privacy policy
4. Ensure proper testing

### Play Store Rejection (Android)

**Symptoms:**
App rejected by Google Play

**Solutions:**
1. Review Play Console policies
2. Check target SDK level
3. Verify app permissions
4. Ensure proper content rating

## Debugging Tools

### iOS
- **Xcode View Debugger**: Visualize view hierarchy
- **Instruments**: Performance analysis
- **Console**: Device logs
- **LLDB**: Debugger command line

### Android
- **Layout Inspector**: Visualize layout
- **Android Profiler**: Performance analysis
- **Logcat**: Application logs
- **Stetho**: Debug bridge

## Getting Help

1. **Check logs**: Xcode console, Logcat
2. **Reproduce**: Try to reproduce consistently
3. **Search**: Stack Overflow, GitHub issues
4. **Ask team**: Post in mobile team channel
5. **Document**: Create issue for persistent problems

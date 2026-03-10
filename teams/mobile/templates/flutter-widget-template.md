# Flutter Widget Template

## Description

Standard Flutter widget template with proper state management, theming, and testing support.

## Stateless Widget Template

```dart
import 'package:flutter/material.dart';

/// MyWidget displays a brief description of what it does.
class MyWidget extends StatelessWidget {
  /// Creates a new MyWidget instance.
  const MyWidget({
    super.key,
    required this.title,
    this.subtitle,
    this.onTap,
  });

  /// The main title text to display.
  final String title;

  /// Optional subtitle text.
  final String? subtitle;

  /// Callback when the widget is tapped.
  final VoidCallback? onTap;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return GestureDetector(
      onTap: onTap,
      child: Container(
        padding: const EdgeInsets.all(16.0),
        decoration: BoxDecoration(
          color: theme.colorScheme.surface,
          borderRadius: BorderRadius.circular(8.0),
          boxShadow: [
            BoxShadow(
              color: theme.shadowColor.withOpacity(0.1),
              blurRadius: 8,
              offset: const Offset(0, 2),
            ),
          ],
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              title,
              style: theme.textTheme.titleLarge,
            ),
            if (subtitle != null) ...[
              const SizedBox(height: 8.0),
              Text(
                subtitle!,
                style: theme.textTheme.bodyMedium?.copyWith(
                  color: theme.colorScheme.onSurface.withOpacity(0.6),
                ),
              ),
            ],
          ],
        ),
      ),
    );
  }
}
```

## Stateful Widget Template

```dart
import 'package:flutter/material.dart';

/// MyCounterWidget demonstrates stateful widget usage.
class MyCounterWidget extends StatefulWidget {
  /// Creates a new MyCounterWidget instance.
  const MyCounterWidget({super.key});

  @override
  State<MyCounterWidget> createState() => _MyCounterWidgetState();
}

class _MyCounterWidgetState extends State<MyCounterWidget> {
  int _counter = 0;

  void _increment() {
    setState(() {
      _counter++;
    });
  }

  void _decrement() {
    setState(() {
      _counter--;
    });
  }

  void _reset() {
    setState(() {
      _counter = 0;
    });
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              'Counter: $_counter',
              style: theme.textTheme.headlineMedium,
            ),
            const SizedBox(height: 16.0),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                IconButton.filled(
                  onPressed: _decrement,
                  icon: const Icon(Icons.remove),
                ),
                const SizedBox(width: 16.0),
                IconButton.filled(
                  onPressed: _increment,
                  icon: const Icon(Icons.add),
                ),
                const SizedBox(width: 16.0),
                IconButton.outlined(
                  onPressed: _reset,
                  icon: const Icon(Icons.refresh),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
```

## Riverpod State Management Template

```dart
import 'package:flutter_riverpod/flutter_riverpod.dart';

// State provider for counter
final counterProvider = StateNotifierProvider<CounterNotifier, int>((ref) {
  return CounterNotifier();
});

// State notifier
class CounterNotifier extends StateNotifier<int> {
  CounterNotifier() : super(0);

  void increment() => state++;
  void decrement() => state--;
  void reset() => state = 0;
}

// Provider for async data
final dataProvider = FutureProvider.autoDispose<Data>((ref) async {
  final response = await fetchData();
  return response;
});

// Widget using Riverpod
class CounterWidget extends ConsumerWidget {
  const CounterWidget({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final counter = ref.watch(counterProvider);
    final dataAsync = ref.watch(dataProvider);

    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Count: $counter'),
            const SizedBox(height: 16),
            dataAsync.when(
              data: (data) => Text('Data: ${data.value}'),
              loading: () => const CircularProgressIndicator(),
              error: (err, stack) => Text('Error: $err'),
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => ref.read(counterProvider.notifier).increment(),
        child: const Icon(Icons.add),
      ),
    );
  }
}
```

## Testing Template

```dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  group('MyWidget Tests', () {
    testWidgets('displays title text', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: MyWidget(title: 'Test Title'),
          ),
        ),
      );

      expect(find.text('Test Title'), findsOneWidget);
    });

    testWidgets('does not display subtitle when null', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: MyWidget(title: 'Title'),
          ),
        ),
      );

      expect(find.byType(Text), findsOneWidget);
    });

    testWidgets('calls onTap when tapped', (tester) async {
      var tapped = false;

      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MyWidget(
              title: 'Title',
              onTap: () => tapped = true,
            ),
          ),
        ),
      );

      await tester.tap(find.byType(GestureDetector));
      expect(tapped, isTrue);
    });
  });
}
```

## Tags

`mobile`, `flutter`, `dart`, `widget`, `template`, `state-management`

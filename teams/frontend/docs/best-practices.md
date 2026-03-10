# Frontend Team Best Practices

## Overview
This document outlines the best practices for frontend development at our company.

## Component Design

### Composition over Inheritance
Build complex UIs by composing simple components together rather than using inheritance.

```tsx
// Good: Composition
<Card>
  <CardHeader>
    <CardTitle>Title</CardTitle>
  </CardHeader>
  <CardBody>Content</CardBody>
  <CardFooter>
    <Button>Action</Button>
  </CardFooter>
</Card>

// Avoid: Deep inheritance hierarchies
```

### Single Responsibility
Each component should have one clear purpose.

```tsx
// Good: Component does one thing
export const UserAvatar = ({ src, alt }: AvatarProps) => (
  <img src={src} alt={alt} className="avatar" />
);

// Avoid: Component doing too many things
export const UserWidget = () => { /* avatar, profile, settings, etc */ }
```

## State Management

### Local State First
Use local state (useState) before reaching for global state solutions.

### State Colocation
Keep state as close to where it's used as possible.

### Immutability
Always update state immutably.

```tsx
// Good
setUsers([...users, newUser]);
setUser({ ...user, name: 'New Name' });

// Bad
setUsers.push(newUser);
user.name = 'New Name';
setUser(user);
```

## Performance

### Code Splitting
Split code by route to reduce initial bundle size.

```tsx
const Dashboard = lazy(() => import('./pages/Dashboard'));
```

### Memoization
Use memo, useMemo, and useCallback judiciously.

```tsx
// ExpensiveList only re-renders when items change
const ExpensiveList = memo(({ items }) => {
  return items.map(item => <Item key={item.id} {...item} />);
});
```

### Virtualization
Use react-window or react-virtualized for long lists.

## Accessibility

### Semantic HTML
Use proper HTML elements.

```tsx
// Good
<button onClick={handleClick}>Click me</button>

// Bad
<div onClick={handleClick}>Click me</div>
```

### ARIA Labels
Provide accessible labels for interactive elements.

```tsx
<button aria-label="Close dialog" onClick={onClose}>
  <XIcon />
</button>
```

### Keyboard Navigation
Ensure all functionality is accessible via keyboard.

## Testing

### Testing Pyramid
- Unit tests: Individual functions and hooks
- Integration tests: Component interactions
- E2E tests: Critical user flows

### Test Isolation
Each test should be independent and able to run in any order.

```tsx
// Good: Isolated test
describe('Button', () => {
  it('renders with label', () => {
    render(<Button>Click me</Button>);
    expect(screen.getByText('Click me')).toBeInTheDocument();
  });

  beforeEach(() => {
    // Reset state between tests
  });
});
```

## Code Style

### Descriptive Names
Use clear, descriptive names for variables, functions, and components.

```tsx
// Good
const getUserById = (id: string) => { /* ... */ };
const UserList = () => { /* ... */ };

// Bad
const getData = (i: string) => { /* ... */ };
const List = () => { /* ... */ };
```

### Early Returns
Use early returns to reduce nesting.

```tsx
// Good
const Component = ({ data, loading, error }) => {
  if (loading) return <Spinner />;
  if (error) return <Error message={error} />;
  if (!data) return <Empty />;

  return <DataView data={data} />;
};

// Less readable
const Component = ({ data, loading, error }) => {
  if (!loading && !error && data) {
    return <DataView data={data} />;
  } else if (loading) {
    return <Spinner />;
  }
  // ...
};
```

## Error Handling

### Error Boundaries
Use error boundaries to catch and handle component errors.

```tsx
<ErrorBoundary fallback={<ErrorFallback />}>
  <App />
</ErrorBoundary>
```

### Loading States
Always provide loading feedback for async operations.

### User-Friendly Messages
Display clear, actionable error messages to users.

## Security

### Input Sanitization
Sanitize user input to prevent XSS attacks.

```tsx
// Good: React escapes by default
<div>{userInput}</div>

// Bad: Dangerous
<div dangerouslySetInnerHTML={{ __html: userInput }} />
```

### Authentication
Never store tokens in localStorage. Use secure, httpOnly cookies.

### API Keys
Never expose API keys in client-side code.

## Resources

- [React Documentation](https://react.dev)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)
- [Web.dev Best Practices](https://web.dev/best-practices/)

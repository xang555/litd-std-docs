# Frontend Team Troubleshooting Guide

## Common Issues and Solutions

## Build Issues

### "Module not found" Error

**Symptoms:**
```
Module not found: Can't resolve './Component'
```

**Solutions:**
1. Check file path is correct (case-sensitive)
2. Ensure file has proper extension (.tsx, .ts)
3. Check if index.ts export exists
4. Clear node_modules and reinstall: `rm -rf node_modules && npm install`

### Type Errors After Update

**Symptoms:**
TypeScript errors after updating dependencies

**Solutions:**
1. Update TypeScript types: `npm install --save-dev @types/react @types/react-dom`
2. Clear cache: `rm -rf .cache dist`
3. Check for breaking changes in library changelog

## Runtime Issues

### "Too many re-renders" Error

**Symptoms:**
```
Warning: Too many re-renders. React limits the number of renders to prevent an infinite loop.
```

**Solutions:**
1. Check for state updates in render body
2. Move state updates to event handlers or useEffect
3. Use functional updates for derived state

```tsx
// Bad
const Component = () => {
  const [count, setCount] = useState(0);
  setCount(count + 1); // Infinite loop!
  return <div>{count}</div>;
};

// Good
const Component = () => {
  const [count, setCount] = useState(0);
  useEffect(() => {
    setCount(1);
  }, []);
  return <div>{count}</div>;
};
```

### Stale Closure Problem

**Symptoms:**
Component uses old state or prop values

**Solutions:**
1. Use functional state updates
2. Add dependencies to useEffect/useCallback/useMemo
3. Use useRef for values that shouldn't trigger re-renders

```tsx
// Good: Functional update
setCount(prev => prev + 1);

// Good: All dependencies listed
useEffect(() => {
  fetchData(userId);
}, [userId]);
```

## Performance Issues

### Slow Component Rendering

**Symptoms:**
UI is slow or janky

**Diagnosis:**
1. Use React DevTools Profiler to identify slow components
2. Check for expensive operations in render
3. Look for unnecessary re-renders

**Solutions:**
1. Memoize expensive calculations with useMemo
2. Memoize callbacks with useCallback
3. Wrap components in React.memo
4. Use code splitting for large components
5. Implement virtual scrolling for long lists

### Large Bundle Size

**Symptoms:**
Slow initial page load

**Solutions:**
1. Analyze bundle: `npm run build:analyze`
2. Implement code splitting
3. Remove unused dependencies
4. Use tree-shaking friendly imports
5. Enable compression in build config

## State Management Issues

### State Not Updating

**Symptoms:**
Changes to state don't trigger re-render

**Solutions:**
1. Ensure you're using setState or useState
2. Check for object mutation (state must be immutable)
3. Verify state is in correct component (not parent/child)

```tsx
// Bad: Mutation
const [user, setUser] = useState({ name: 'John' });
user.name = 'Jane'; // Won't trigger update
setUser(user);

// Good: New object
setUser({ ...user, name: 'Jane' });
```

### Props Drilling

**Symptoms:**
Passing props through many intermediate components

**Solutions:**
1. Use context for shared state
2. Consider state management library (Redux, Zustand)
3. Component composition

## Styling Issues

### Styles Not Applying

**Symptoms:**
CSS classes not working as expected

**Solutions:**
1. Check CSS specificity
2. Verify CSS is imported
3. Check for typos in class names
4. Inspect element in DevTools
5. Check CSS-in-JS library configuration

### Responsive Layout Issues

**Symptoms:**
Layout breaks on different screen sizes

**Solutions:**
1. Use responsive units (rem, em, %, vw/vh)
2. Implement proper breakpoints
3. Test on multiple devices
4. Check for fixed widths
5. Use CSS Grid or Flexbox

## Testing Issues

### Test Fails Intermittently

**Symptoms:**
Test passes sometimes, fails other times

**Solutions:**
1. Check for race conditions
2. Use proper async/await patterns
3. Add proper cleanup in afterEach
4. Mock external dependencies
5. Increase timeout if needed

```tsx
// Good: Proper async handling
it('loads data', async () => {
  render(<Component />);
  await waitFor(() => {
    expect(screen.getByText('Loaded')).toBeInTheDocument();
  });
});
```

### Mock Not Working

**Symptoms:**
Mocked function not being called

**Solutions:**
1. Check mock implementation
2. Verify mock is cleared between tests
3. Use jest.clearAllMocks() in beforeEach
4. Check for multiple instances of module

## Getting Help

If none of these solutions work:

1. **Check logs**: Browser console, terminal, build logs
2. **Search issues**: GitHub issues, Stack Overflow
3. **Ask team**: Post in frontend team channel
4. **Create issue**: Document the problem for future reference

## Debug Tools

- **React DevTools**: Component tree, props, state
- **Redux DevTools**: State changes, actions
- **Browser DevTools**: Elements, Console, Network
- **Lighthouse**: Performance, accessibility, SEO

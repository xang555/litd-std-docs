# React Conventions Standard

## Scope
This standard applies to all React components and related code in frontend applications.

## Rules

### 1. Component Structure

**Rule:** Components must follow a consistent file structure.

**Rationale:** Consistent structure improves readability and makes code easier to navigate.

**Example:**
```
ComponentName/
├── ComponentName.tsx
├── ComponentName.test.tsx
├── ComponentName.styles.ts
├── ComponentName.types.ts
└── index.ts
```

### 2. Naming Conventions

**Rule:** Use PascalCase for component names, camelCase for props and handlers.

**Rationale:** Follows React community conventions and clearly distinguishes components from variables.

**Example:**
```tsx
// Good
export const UserProfile: React.FC<UserProfileProps> = ({ userName, onUserClick }) => {
  return <div>{userName}</div>;
};

// Bad
export const userProfile = ({ user_name, on_user_click }) => {
  return <div>{user_name}</div>;
};
```

### 3. Prop Types

**Rule:** All components must have defined prop types using TypeScript interfaces.

**Rationale:** Type safety prevents runtime errors and improves developer experience.

**Example:**
```tsx
interface ButtonProps {
  label: string;
  onClick: () => void;
  disabled?: boolean;
  variant?: 'primary' | 'secondary';
}
```

### 4. Hooks Usage

**Rule:** Only call hooks at the top level of React functions.

**Rationale:** Hooks rely on call order; conditional calls break this contract.

**Example:**
```tsx
// Good
const Component = () => {
  const [state, setState] = useState();
  useEffect(() => {}, []);
  return <div />;
};

// Bad
const Component = () => {
  if (condition) {
    const [state, setState] = useState(); // Don't do this
  }
  return <div />;
};
```

### 5. Component Size

**Rule:** Components should be under 250 lines. Split larger components into smaller sub-components.

**Rationale:** Smaller components are easier to understand, test, and maintain.

### 6. Performance

**Rule:** Use React.memo, useMemo, and useCallback for optimization when needed.

**Rationale:** Prevents unnecessary re-renders and improves application performance.

## Enforcement

### Linting
- ESLint configuration: `.eslintrc.json`
- Run: `npm run lint`

### Review
All pull requests must:
- Pass ESLint checks
- Pass TypeScript type checking
- Have component tests with >80% coverage
- Be reviewed by at least one team member

### Pre-commit Hooks
- ESLint runs automatically
- Prettier formats code automatically
- Type checking runs before push

## Exceptions

Exceptions to these standards require:
- Team lead approval
- Documentation of the exception rationale
- Comment in code explaining the deviation

## Related Standards
- TypeScript Conventions
- Testing Standards
- Accessibility Guidelines

## Tags
`frontend`, `react`, `standards`, `conventions`, `typescript`

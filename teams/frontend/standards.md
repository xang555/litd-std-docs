# Frontend Team Standards

## Overview

This document outlines the coding standards, best practices, and conventions for the Frontend Engineering team.

## Technology Stack

- **Framework**: React 18+ with TypeScript
- **Styling**: Tailwind CSS + CSS Modules
- **State Management**: Zustand / React Query
- **Testing**: Vitest + React Testing Library
- **Build Tool**: Vite

## Code Standards

### File Naming

- Components: `PascalCase.tsx` (e.g., `UserProfile.tsx`)
- Utilities: `kebab-case.ts` (e.g., `format-date.ts`)
- Hooks: `use*.ts` (e.g., `useAuth.ts`)
- Types: `*.types.ts` (e.g., `user.types.ts`)

### Component Structure

```tsx
/**
 * Component description
 */
export function ComponentName({ prop }: Props) {
  // 1. Hooks
  // 2. Derived state
  // 3. Effects
  // 4. Handlers
  // 5. Render
}
```

### TypeScript Rules

- Enable strict mode
- Avoid `any` - use `unknown` when type is truly unknown
- Use type inference when possible
- Export types separately: `export type { ComponentProps }`

## Best Practices

- Components should be < 200 lines
- Prefer composition over inheritance
- Use semantic HTML
- Implement proper error boundaries
- All public components must have Storybook stories

## Related Documents

- [React Guidelines](./react-guidelines.md)
- [CSS Architecture](./css-architecture.md)

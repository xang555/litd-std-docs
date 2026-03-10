# React Component Template

## Description

Standard React component template with TypeScript, following best practices for structure, typing, and documentation.

## Usage

Copy this template to create new React components. Replace `ComponentName` with your actual component name.

## File Structure

```
ComponentName/
├── ComponentName.tsx      # Main component file
├── ComponentName.types.ts # TypeScript types
├── ComponentName.styles.ts# Styled components
└── index.ts               # Barrel export
```

## Component Template

```tsx
/**
 * ComponentName
 *
 * Brief description of what this component does.
 *
 * @param props - Component properties
 * @returns JSX element
 */

import { useState } from 'react'

// ============================================================================
// Types
// ============================================================================

interface ComponentNameProps {
  /** Description of prop1 */
  prop1: string
  /** Description of prop2 */
  prop2?: number
  /** Optional callback function */
  onAction?: (value: string) => void
}

// ============================================================================
// Component
// ============================================================================

/**
 * ComponentName component
 */
export function ComponentName({
  prop1,
  prop2 = 0,
  onAction
}: ComponentNameProps) {
  // ------------------------------------------------------------------------
  // State
  // ------------------------------------------------------------------------
  const [localState, setLocalState] = useState<string>('')

  // ------------------------------------------------------------------------
  // Derived State
  // ------------------------------------------------------------------------
  const derivedValue = `${prop1}-${prop2}`

  // ------------------------------------------------------------------------
  // Effects
  // ------------------------------------------------------------------------
  // useEffect(() => {
  //   // Effect logic here
  // }, [dependencies])

  // ------------------------------------------------------------------------
  // Handlers
  // ------------------------------------------------------------------------
  const handleClick = () => {
    onAction?.(localState)
  }

  // ------------------------------------------------------------------------
  // Render
  // ------------------------------------------------------------------------
  return (
    <div className="component-name">
      <h2>{prop1}</h2>
      <p>{derivedValue}</p>
      <button onClick={handleClick}>
        Action
      </button>
    </div>
  )
}

// ============================================================================
// Default Export
// ============================================================================

export default ComponentName
```

## Types Template

```tsx
export interface ComponentNameProps {
  /**
   * Component content
   */
  children: React.ReactNode;

  /**
   * Visual variant
   */
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost';

  /**
   * Component size
   */
  size?: 'small' | 'medium' | 'large';

  /**
   * Disable the component
   */
  disabled?: boolean;

  /**
   * Show loading state
   */
  loading?: boolean;

  /**
   * Click handler
   */
  onClick?: () => void;

  /**
   * Additional CSS class
   */
  className?: string;
}
```

## Styles Template (Styled Components)

```tsx
import styled, { css } from 'styled-components';

type ComponentVariant = 'primary' | 'secondary' | 'outline' | 'ghost';
type ComponentSize = 'small' | 'medium' | 'large';

interface StyledComponentProps {
  $variant: ComponentVariant;
  $size: ComponentSize;
  disabled: boolean;
}

const sizeStyles = {
  small: css`
    padding: 0.375rem 0.75rem;
    font-size: 0.875rem;
  `,
  medium: css`
    padding: 0.5rem 1rem;
    font-size: 1rem;
  `,
  large: css`
    padding: 0.75rem 1.5rem;
    font-size: 1.125rem;
  `,
};

const variantStyles = {
  primary: css`
    background-color: #007bff;
    color: white;
    &:hover:not(:disabled) {
      background-color: #0056b3;
    }
  `,
  secondary: css`
    background-color: #6c757d;
    color: white;
    &:hover:not(:disabled) {
      background-color: #545b62;
    }
  `,
  outline: css`
    background-color: transparent;
    border: 1px solid #007bff;
    color: #007bff;
    &:hover:not(:disabled) {
      background-color: #007bff;
      color: white;
    }
  `,
  ghost: css`
    background-color: transparent;
    color: #007bff;
    &:hover:not(:disabled) {
      background-color: rgba(0, 123, 255, 0.1);
    }
  `,
};

export const StyledComponent = styled.button<StyledComponentProps>`
  border: none;
  border-radius: 0.375rem;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease-in-out;

  ${({ $size }) => sizeStyles[$size]}
  ${({ $variant }) => variantStyles[$variant]}

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
`;
```

## Index Export Template

```tsx
export { ComponentName } from './ComponentName';
export type { ComponentNameProps } from './ComponentName.types';
```

## Best Practices

1. **Named Exports**: Use named exports for components
2. **Type Safety**: Define props interfaces in separate `.types.ts` files
3. **Default Props**: Use destructuring with default values
4. **Comments**: Use JSDoc comments for component documentation
5. **Structure**: Separate state, effects, handlers, and render logic

## Tags

`frontend`, `react`, `component`, `template`, `typescript`

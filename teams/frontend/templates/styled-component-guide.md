# Styled Component Guide

## Description

Guide and templates for creating styled components with proper variant and size support.

## Pattern Overview

This pattern uses styled-components with polymorphic props for creating reusable UI components with multiple variants and sizes.

## Styled Component Template

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
    border-radius: 0.25rem;
  `,
  medium: css`
    padding: 0.5rem 1rem;
    font-size: 1rem;
    border-radius: 0.375rem;
  `,
  large: css`
    padding: 0.75rem 1.5rem;
    font-size: 1.125rem;
    border-radius: 0.5rem;
  `,
};

const variantStyles = {
  primary: css`
    background-color: #007bff;
    color: white;
    &:hover:not(:disabled) {
      background-color: #0056b3;
    }
    &:active:not(:disabled) {
      background-color: #004085;
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

export const StyledButton = styled.button<StyledButtonProps>`
  border: none;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease-in-out;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;

  ${({ $size }) => sizeStyles[$size]}
  ${({ $variant }) => variantStyles[$variant]}

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  &:focus-visible {
    outline: 2px solid #007bff;
    outline-offset: 2px;
  }
`;
```

## Component with Styled Props

```tsx
import React from 'react';
import { StyledButton } from './Button.styles';

export const Button: React.FC<ButtonProps> = ({
  children,
  variant = 'primary',
  size = 'medium',
  disabled = false,
  loading = false,
  onClick,
  className,
  ...props
}) => {
  return (
    <StyledButton
      $variant={variant}
      $size={size}
      disabled={disabled || loading}
      onClick={onClick}
      className={className}
      {...props}
    >
      {loading ? (
        <>
          <Spinner size="small" />
          Loading...
        </>
      ) : children}
    </StyledButton>
  );
};
```

## Theme Integration

```tsx
import styled, { css } from 'styled-components';
import { theme } from './theme';

const variantStyles = {
  primary: css`
    background-color: ${props => props.theme.colors.primary};
    color: ${props => props.theme.colors.onPrimary};
  `,
  // ... other variants
};
```

## Animation Template

```tsx
import styled, { keyframes, css } from 'styled-components';

const spin = keyframes`
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
`;

export const Spinner = styled.div<{ $size: number }>`
  animation: ${spin} 1s linear infinite;
  border: 2px solid transparent;
  border-top-color: currentColor;
  border-radius: 50%;

  ${({ $size }) => css`
    width: ${$size}px;
    height: ${$size}px;
  `}
`;
```

## Responsive Styles

```tsx
import styled, { css } from 'styled-components';

const responsive = css`
  padding: 1rem;

  @media (max-width: 768px) {
    padding: 0.75rem;
  }

  @media (max-width: 480px) {
    padding: 0.5rem;
  }
`;
```

## Tags

`frontend`, `styled-components`, `css`, `styling`, `template`

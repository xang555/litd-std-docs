# Component Create Command

## Description
Command template for scaffolding new React components with best practices.

## Usage
```
/component-create <ComponentName> [options]
```

## Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| ComponentName | string | true | - | Name of the component (PascalCase) |
| --type | string | false | functional | Component type: functional, class |
| --with-styles | boolean | false | true | Include styled component |
| --with-test | boolean | false | true | Include test file |
| --with-types | boolean | false | true | Include TypeScript types |
| --path | string | false | components | Target directory path |

## Examples

### Basic component
```bash
/component-create Button
```
Creates: `components/Button/Button.tsx`

### Component with options
```bash
/component-create DataTable --with-styles=false --path=components/data
```
Creates: `components/data/DataTable/DataTable.tsx`

### Component in specific location
```bash
/component-create UserAvatar --path=components/users
```
Creates: `components/users/UserAvatar/UserAvatar.tsx`

## Generated Files

When run with default options, generates:
```
ComponentName/
├── ComponentName.tsx       # Component code
├── ComponentName.test.tsx  # Test file
├── ComponentName.styles.ts # Styled components
├── index.ts                # Export file
└── ComponentName.types.ts  # TypeScript types
```

## Component Template

The generated component follows this structure:

```tsx
import React from 'react';
import { ComponentNameProps } from './ComponentName.types';
import { StyledComponentName } from './ComponentName.styles';

export const ComponentName: React.FC<ComponentNameProps> = ({
  prop1,
  prop2,
}) => {
  return (
    <StyledComponentName>
      {/* Component content */}
    </StyledComponentName>
  );
};

export default ComponentName;
```

## Tags
`frontend`, `react`, `component`, `scaffold`, `cli`

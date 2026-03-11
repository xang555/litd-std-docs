---
name: feature
description: Expert in feature-first architecture, project structure, and file organization for modular CRUD features
---
# Skill: Feature-First Architecture (2026)

## Description

Expert in **feature-first architecture** and **project structure** for building modular, self-contained CRUD features. Focuses on proper file organization, folder structure, naming conventions, and architectural patterns rather than implementation details.

**See Examples:** `examples/` folder contains complete feature structure with all preset files.

## Core Architecture

### Feature-First Structure

**Philosophy:** Each feature is a self-contained module with all its business logic, UI components, data fetching, and types in one place.

### Complete Directory Structure

```
src/features/{feature}/
├── components/              # UI Components
│   ├── {Feature}List.tsx       # List/table component
│   ├── {Feature}Filters.tsx    # Filter accordion component
│   ├── {Feature}Form.tsx       # Form dialog/page component
│   └── index.ts                # Export all components
├── hooks/                   # Data Fetching Hooks
│   ├── use{Features}.ts        # List query hook
│   ├── use{Feature}.ts         # Detail query hook
│   ├── useCreate{Feature}.ts   # Create mutation hook
│   ├── useUpdate{Feature}.ts   # Update mutation hook
│   ├── useDelete{Feature}.ts   # Delete mutation hook (optional)
│   └── index.ts                # Export all hooks
├── schemas/                 # Validation Schemas
│   └── {feature}.schema.ts     # Zod schemas (as functions)
├── services/                # API Service Layer
│   └── {feature}.service.ts    # HTTP requests
├── types/                   # TypeScript Definitions
│   └── index.ts                # Interfaces and types
├── constants/               # Static Configuration
│   └── {feature}.constants.ts  # Endpoints, query keys, enums
└── index.ts                 # Public API (barrel export)
```

### Route Structure

```
src/routes/_layout/{feature}/
└── index.tsx                # Route + Page component + Search schema
```

## Architectural Layers

### 1. Components Layer (`components/`)

**Purpose:** UI presentation and user interaction

**Files:**
- `{Feature}List.tsx` - List/table view
- `{Feature}Filters.tsx` - Filter controls
- `{Feature}Form.tsx` - Create/edit form
- `index.ts` - Barrel export

**Responsibilities:**
- Render UI components
- Handle user interactions
- Use hooks for data
- No direct API calls
- No business logic

**See:** Form and List-Table skills for implementation

### 2. Hooks Layer (`hooks/`)

**Purpose:** Data fetching and state management

**Files:**
- `use{Features}.ts` - List query (GET /items)
- `use{Feature}.ts` - Detail query (GET /items/:id)
- `useCreate{Feature}.ts` - Create mutation (POST /items)
- `useUpdate{Feature}.ts` - Update mutation (PUT /items/:id)
- `useDelete{Feature}.ts` - Delete mutation (DELETE /items/:id)
- `index.ts` - Barrel export

**Responsibilities:**
- Wrap TanStack Query hooks
- Handle cache invalidation
- Provide loading/error states
- No UI rendering
- No direct axios calls

### 3. Services Layer (`services/`)

**Purpose:** HTTP communication with backend API

**Files:**
- `{feature}.service.ts` - API methods

**Responsibilities:**
- Make HTTP requests via axiosInstance
- Accept AbortSignal for cancellation
- Return typed responses
- No state management
- No UI logic

### 4. Schemas Layer (`schemas/`)

**Purpose:** Data validation and type generation

**Files:**
- `{feature}.schema.ts` - Zod schemas

**Responsibilities:**
- Define validation rules
- Accept TFunction for i18n
- Export as functions
- Generate TypeScript types
- No API calls

### 5. Types Layer (`types/`)

**Purpose:** TypeScript type definitions

**Files:**
- `index.ts` - Interfaces and types

**Responsibilities:**
- Define data structures
- Export API response types
- Re-export schema types
- No implementation

### 6. Constants Layer (`constants/`)

**Purpose:** Static configuration values

**Files:**
- `{feature}.constants.ts` - Endpoints, query keys, enums

**Responsibilities:**
- API endpoint URLs
- Query key hierarchies
- Enum values
- Default values
- No logic

### 7. Route Layer (`routes/_layout/{feature}/`)

**Purpose:** URL routing and page orchestration

**Files:**
- `index.tsx` - Route definition + Page component

**Responsibilities:**
- Define search params schema
- Validate URL parameters
- Orchestrate components
- Manage page-level state
- Handle navigation

### 8. Translation Layer (`locales/{lang}/`)

**Purpose:** Internationalization support

**Files:**
- `en/translation.json` - English translations
- `lo/translation.json` - Lao translations

**Structure:**
```json
{
  "{feature}": {
    "list": "Title",
    "add": "Add New",
    "validation": {
      "fieldRequired": "Field is required"
    }
  }
}
```

---

## Menu Integration & Routing

### Adding Feature to Navigation Menu

**Common Menu Patterns:**
1. **Sidebar Menu Item** - Single menu entry with icon and path
2. **Grouped Menu Items** - Nested menu with parent and children
3. **Permission-Based Display** - Show/hide based on user permissions

**Menu Configuration:**
- Label with translation key
- Icon component
- Path to feature route
- Optional permission check
- Optional nested children

### Route Registration

**File-Based Routing:**
- Routes auto-register based on file structure
- File location = URL path
- No manual registration needed

**Route Structure Pattern:**
```
src/routes/_layout/{feature}/
├── index.tsx          # /{feature} (list page)
├── $id.tsx            # /{feature}/:id (detail page)
└── $id.edit.tsx       # /{feature}/:id/edit (edit page)
```

### Breadcrumb Integration

**In Route Definition:**
- Add `breadcrumb` key in `beforeLoad()`
- Use translation key for label
- Breadcrumb component auto-renders trail

**Pattern:**
```
beforeLoad: () => ({ breadcrumb: 'breadcrumbs.featureName' })
```

### Navigation Patterns

**Common Actions:**
- List → Detail (view item)
- Detail → List (back button)
- List → Edit (preserve filters)
- After Create → Detail or stay on list
- After Update → Stay (auto-refresh via cache)
- After Delete → Stay (auto-refresh via cache)

**Search Params:**
- Preserve filters when navigating
- Reset to page 1 when filtering
- Use `search: (prev) => prev` to keep params

### Permission-Based Access

**Menu Level:**
- Check permissions before rendering menu items
- Hide items user cannot access
- Support nested permission checks

**Route Level:**
- Optional permission check in `beforeLoad()`
- Redirect to unauthorized page if needed

### Translation Keys Structure

**Required Keys:**
- `menu.{feature}` - Menu label
- `breadcrumbs.{feature}` - Breadcrumb label
- `breadcrumbs.{feature}Detail` - Detail page label

**See:** `examples/menu-and-routing.md` for complete implementation

## Naming Conventions

### Files and Folders
- **Feature folder**: lowercase, singular (e.g., `position`, `user`, `vehicle`)
- **Component files**: PascalCase (e.g., `PositionList.tsx`, `UserForm.tsx`)
- **Hook files**: camelCase with `use` prefix (e.g., `usePositions.ts`, `useCreateUser.ts`)
- **Service files**: lowercase with `.service` suffix (e.g., `position.service.ts`)
- **Schema files**: lowercase with `.schema` suffix (e.g., `user.schema.ts`)
- **Type files**: `index.ts` in types folder
- **Constants files**: lowercase with `.constants` suffix (e.g., `vehicle.constants.ts`)

### Exports
- **Barrel exports**: Use `index.ts` in each folder
- **Named exports**: Always use named exports, never default
- **Public API**: Feature root `index.ts` exports only public interfaces

### TypeScript
- **Interfaces**: PascalCase (e.g., `Position`, `User`, `VehicleFormData`)
- **Types**: PascalCase (e.g., `PositionsResponse`, `UserSearch`)
- **Constants**: UPPER_SNAKE_CASE (e.g., `POSITION_ENDPOINTS`, `USER_QUERY_KEYS`)
- **Functions**: camelCase (e.g., `getPositionSchema`, `createUser`)

## Architectural Rules

### Structure Rules

1. **Feature Encapsulation** - All business logic stays within feature folder
2. **Layer Separation** - Components don't call services directly
3. **Single Responsibility** - Each file has one clear purpose
4. **Barrel Exports** - Use `index.ts` for clean imports
5. **Public API** - Feature root exports only public interfaces
6. **No Cross-Feature Imports** - Features don't import from each other
7. **Shared for Generic Only** - Only truly generic code in `src/shared/`

### File Organization Rules

1. **Consistent Naming** - Follow naming conventions strictly
2. **Folder Structure** - Match the standard structure exactly
3. **File Placement** - Put files in correct layer folders
4. **Index Files** - Every folder with multiple files needs `index.ts`
5. **Route Colocation** - Route file contains page component

### Dependency Rules

1. **Unidirectional Flow**:
   ```
   Components → Hooks → Services → API
        ↓
     Schemas
        ↓
      Types
   ```
2. **No Circular Dependencies** - Imports flow one direction
3. **Types First** - Types/interfaces defined before implementation
4. **Constants Shared** - Query keys, endpoints in constants file

## File Templates

### Constants File Structure
```typescript
// API Endpoints
export const {NAME}_ENDPOINTS = {
  LIST: '/{names}',
  DETAIL: '/{names}/:id',
  CREATE: '/{names}',
  UPDATE: '/{names}/:id',
  DELETE: '/{names}/:id',
} as const;

// Query Keys (Hierarchical)
export const {NAME}_QUERY_KEYS = {
  all: ['{names}'] as const,
  lists: () => [...{NAME}_QUERY_KEYS.all, 'list'] as const,
  list: (params: any) => [...{NAME}_QUERY_KEYS.lists(), params] as const,
  details: () => [...{NAME}_QUERY_KEYS.all, 'detail'] as const,
  detail: (id: number) => [...{NAME}_QUERY_KEYS.details(), id] as const,
} as const;
```

### Service File Structure
```typescript
import axiosInstance from '@/core/api/axios.config';
import { {NAME}_ENDPOINTS } from '../constants/{name}.constants';
import type { {Name}, {Names}Response } from '../types';

export const {name}Service = {
  async getAll(params: any, signal?: AbortSignal): Promise<{Names}Response> { },
  async getById(id: number, signal?: AbortSignal): Promise<{ {name}: {Name} }> { },
  async create(data: any, signal?: AbortSignal): Promise<{ {name}: {Name} }> { },
  async update(id: number, data: any, signal?: AbortSignal): Promise<{ {name}: {Name} }> { },
  async delete(id: number, signal?: AbortSignal): Promise<void> { },
};
```

### Hook File Structure
```typescript
// List Query Hook
import { useQuery } from '@tanstack/react-query';
import { {name}Service } from '../services/{name}.service';
import { {NAME}_QUERY_KEYS } from '../constants/{name}.constants';

export function use{Names}(params: any) {
  return useQuery({
    queryKey: {NAME}_QUERY_KEYS.list(params),
    queryFn: ({ signal }) => {name}Service.getAll(params, signal),
    staleTime: 5 * 60 * 1000,
  });
}

// Mutation Hook
import { useMutation, useQueryClient } from '@tanstack/react-query';

export function useCreate{Name}() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: any) => {name}Service.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: {NAME}_QUERY_KEYS.lists() });
    },
  });
}
```

### Schema File Structure
```typescript
import { z } from 'zod';
import type { TFunction } from 'i18next';

export const get{Name}Schema = (t: TFunction) => z.object({
  name: z.string().min(1, t('{names}.validation.nameRequired')),
});

export type {Name}FormData = z.infer<ReturnType<typeof get{Name}Schema>>;
```

## Feature Checklist

When creating a new feature, ensure:
### Structure
- [ ] Feature folder created in `src/features/{name}/`
- [ ] All 7 subfolders created (components, hooks, schemas, services, types, constants)
- [ ] Route file created in `src/routes/_layout/{name}/index.tsx`
- [ ] Barrel exports (`index.ts`) in each subfolder
- [ ] Public API barrel export at feature root

### Files
- [ ] Constants file with endpoints and query keys
- [ ] Service file with all CRUD methods
- [ ] Schema file as function accepting TFunction
- [ ] Types file with interfaces
- [ ] Hook files for queries and mutations
- [ ] Component files (List, Filters, Form)
- [ ] Route file with search schema

### Translations
- [ ] English translations added
- [ ] Lao translations added
- [ ] Validation messages translated
- [ ] UI labels translated

### Integration
- [ ] Feature exported from root index.ts
- [ ] Route registered in router
- [ ] Navigation links added (if needed)

## Related Documentation

**For Implementation Details:**
- Form implementation → `/.windsurf/skills/form/SKILL.md`
- List/table implementation → `/.windsurf/skills/list-table/SKILL.md`

**For Step-by-Step Guides:**
- Create feature → `/.windsurf/workflows/create-feature.md`
- Implement form → `/.windsurf/workflows/form-implementation.md`
- Implement list → `/.windsurf/workflows/list-table-implementation.md`

**Reference Implementation:**
- `src/features/positions/` - Complete example feature

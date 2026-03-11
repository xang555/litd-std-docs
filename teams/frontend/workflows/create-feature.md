---
description: Step-by-step guide for creating a complete feature module with proper structure and organization
---
# Workflow: Create Complete Feature Module

## Overview

Step-by-step guide for scaffolding a complete CRUD feature following feature-first architecture. Focuses on creating the proper file structure and preset files before implementation.

## Prerequisites

- Feature name in singular form (e.g., "position", "user", "vehicle")
- API endpoints documented
- Database schema known
- Review `/.windsurf/skills/feature/SKILL.md` for architecture concepts
- Check `/.windsurf/skills/feature/examples/` for complete examples

---

## Decision Tree: Plan Your Feature

### 1️⃣ What type of feature are you building?

**Simple CRUD** → Standard structure (most common)
- List page with filters
- Create/Edit form
- Basic CRUD operations

**Read-Only** → Skip form and mutations
- List page only
- Detail page (optional)
- No create/update/delete

**Complex Feature** → Extended structure
- Multiple related entities
- Additional custom components
- Complex business logic

### 2️⃣ Do you need filters?

**Yes** → Create filter component
- Search functionality
- Category/status filters
- Date range filters

**No** → Skip filter component
- Simple list only
- Pagination only

### 3️⃣ What form layout do you need?

**Dialog Form** → < 10 fields
- Opens in modal
- Quick create/edit

**Full Page Form** → 10+ fields
- Dedicated page
- Complex forms with sections

**No Form** → Read-only feature
- Skip form component
- Skip create/update mutations

### 4️⃣ Do you need detail page?

**Yes** → Create detail route
- View full item details
- Additional route file

**No** → List only
- Edit in dialog/form page

---

## Quick Start Paths

### Path A: Simple CRUD (Most Common)
```
Step 1 → Step 2 → Step 3 → Step 4 → Step 5 → Step 6 → Done
```
**Time:** ~30-45 minutes (structure only)  
**Example:** Products, Categories, Tags

### Path B: Read-Only Feature
```
Step 1 → Step 2 (skip mutations) → Step 3 → Step 6 → Done
```
**Time:** ~20 minutes  
**Example:** Reports, Logs, Analytics

### Path C: Complex Feature
```
Step 1 → Step 2 → Step 3 → Step 4 → Step 5 → Add custom components → Step 6 → Done
```
**Time:** ~1 hour  
**Example:** Vehicles with sections, Users with roles

---

## Step 1: Create Directory Structure

### Location
`src/features/{name}/`

### Tasks
1. ✅ Create feature folder (lowercase, singular)
2. ✅ Create 6 subfolders: `components/`, `hooks/`, `schemas/`, `services/`, `types/`, `constants/`
3. ✅ Create `index.ts` in each subfolder (barrel exports)
4. ✅ Create `index.ts` at feature root (public API)

### Folder Structure
```
src/features/{name}/
├── components/
├── hooks/
├── schemas/
├── services/
├── types/
├── constants/
└── index.ts
```

### Reference
**See:** `/.windsurf/skills/feature/SKILL.md` → Complete Directory Structure

---

## Step 2: Create Constants File

### Location
`src/features/{name}/constants/{name}.constants.ts`

### Tasks
1. ✅ Define API endpoints (LIST, DETAIL, CREATE, UPDATE, DELETE)
2. ✅ Define hierarchical query keys
3. ✅ Add enums (if needed)
4. ✅ Use UPPER_SNAKE_CASE for constants

### Key Points
- Endpoints use REST patterns
- Query keys are hierarchical for cache management
- Use `as const` for type safety

### Reference
**See:** `/.windsurf/skills/feature/examples/complete-feature-structure.md` → Constants File

---

## Step 3: Create Service File

### Location
`src/features/{name}/services/{name}.service.ts`

### Tasks
1. ✅ Import axiosInstance from core
2. ✅ Import endpoints from constants
3. ✅ Import types
4. ✅ Create service object with CRUD methods
5. ✅ All methods accept `signal?: AbortSignal`
6. ✅ Return typed responses

### Methods to Create
- `getAll(params, signal)` - GET list
- `getById(id, signal)` - GET single item
- `create(data, signal)` - POST create
- `update(id, data, signal)` - PUT update
- `delete(id, signal)` - DELETE (optional)

### Key Points
- Use axiosInstance, not fetch
- Always accept and pass signal parameter
- Replace `:id` in endpoint URLs
- Return `response.data`

### Reference
**See:** `/.windsurf/skills/feature/examples/complete-feature-structure.md` → Service File

---

## Step 4: Create Types File

### Location
`src/features/{name}/types/index.ts`

### Tasks
1. ✅ Define main entity interface
2. ✅ Define API response interfaces
3. ✅ Re-export form data type from schema

### Interfaces to Create
- `{Name}` - Main entity (id, name, created_at, etc.)
- `{Names}Response` - List response ({names}[], total)
- Re-export `{Name}FormData` from schema

### Key Points
- Use PascalCase for interfaces
- Match API response structure
- Include all fields from backend

### Reference
**See:** `/.windsurf/skills/feature/examples/complete-feature-structure.md` → Types File

---

## Step 5: Create Schema File

### Location
`src/features/{name}/schemas/{name}.schema.ts`

### Tasks
1. ✅ Import Zod and TFunction
2. ✅ Create schema as function accepting TFunction
3. ✅ Add validation rules for each field
4. ✅ Use translation keys for error messages
5. ✅ Export inferred type

### Schema Pattern
- Function name: `get{Name}Schema`
- Parameter: `(t: TFunction)`
- Return: `z.object({ ... })`
- Export type: `z.infer<ReturnType<typeof get{Name}Schema>>`

### Key Points
- Schema is a function, not a constant
- All validation messages use `t()` function
- Translation keys: `{names}.validation.{field}{Rule}`

### Reference
**See:** `/.windsurf/skills/feature/examples/complete-feature-structure.md` → Schema File

---

## Step 6: Create Hook Files

### 6.1 List Query Hook

**File**: `src/features/{name}/hooks/use{Names}.ts`

**Tasks:**
1. ✅ Import useQuery, service, query keys
2. ✅ Accept search params as parameter
3. ✅ Use hierarchical query key with params
4. ✅ Pass signal to service
5. ✅ Configure staleTime (5 minutes)

**Key Points:**
- Query key includes params for cache differentiation
- Signal enables automatic cancellation
- StaleTime prevents unnecessary refetches

### 6.2 Detail Query Hook

**File**: `src/features/{name}/hooks/use{Name}.ts`

**Tasks:**
1. ✅ Import useQuery, service, query keys
2. ✅ Accept id and optional options
3. ✅ Use detail query key with id
4. ✅ Add enabled condition (id > 0)
5. ✅ Pass signal to service

**Key Points:**
- Enabled prevents invalid queries
- Options allow flexibility
- Signal for cancellation

### 6.3 Create Mutation Hook

**File**: `src/features/{name}/hooks/useCreate{Name}.ts`

**Tasks:**
1. ✅ Import useMutation, useQueryClient
2. ✅ Import service and query keys
3. ✅ MutationFn calls service.create
4. ✅ OnSuccess invalidates list queries
5. ✅ NO try-catch in mutationFn

**Key Points:**
- Simple pass-through to service
- Invalidate lists() query key
- Errors handled by onError callback

### 6.4 Update Mutation Hook

**File**: `src/features/{name}/hooks/useUpdate{Name}.ts`

**Tasks:**
1. ✅ Accept id as hook parameter
2. ✅ MutationFn calls service.update(id, data)
3. ✅ OnSuccess invalidates all queries

**Key Points:**
- ID as parameter for cleaner usage
- Invalidate broader scope (all)
- Component: `mutate(formData)` not `mutate({ id, data })`

### 6.5 Delete Mutation Hook (Optional)

**File**: `src/features/{name}/hooks/useDelete{Name}.ts`

**Tasks:**
1. ✅ MutationFn accepts id
2. ✅ Calls service.delete(id)
3. ✅ Invalidates list queries

### 6.6 Export Hooks

**File**: `src/features/{name}/hooks/index.ts`

**Tasks:**
1. ✅ Export all hook files
2. ✅ Use named exports

### Reference
**See:** `/.windsurf/skills/feature/examples/complete-feature-structure.md` → Hook Files

---

## Step 7: Create Route File

Create a route file at `src/routes/_layout/{names}/index.tsx` with the following tasks:

1. Define search params schema with Zod
2. Add pagination params (page, page_size)
3. Add filter params (optional)
4. Export search type
5. Create route with zodValidator
6. Add breadcrumb in beforeLoad
7. Create page component

### Reference
**See:** `/.windsurf/skills/feature/examples/page-structure.md`

---

## Step 8: Create Component Files

Create component files at `src/features/{name}/components/` with the following tasks:

**File**: `src/features/{name}/components/{Name}List.tsx`

**Tasks:**
1. ✅ Define props interface
2. ✅ Fetch data with query hook
3. ✅ Define columns with useMemo
4. ✅ Add row actions dropdown
5. ✅ Render DataTable

**Reference:** `/.windsurf/workflows/list-table-implementation.md`

### 8.2 Filter Component (Optional)

**File**: `src/features/{name}/components/{Name}Filters.tsx`

**Tasks:**
1. ✅ Local state for filter values
2. ✅ Sync with URL using useEffect
3. ✅ Apply button (reset to page 1)
4. ✅ Reset button
5. ✅ Accordion layout

**Reference:** `/.windsurf/workflows/list-table-implementation.md` → Step 3

### 8.3 Form Component

**File**: `src/features/{name}/components/{Name}Form.tsx`

**Tasks:**
1. ✅ Use schema function with t
2. ✅ Setup form with useAppForm
3. ✅ Add form fields
4. ✅ Handle submit with mutation
5. ✅ Dialog or full-page layout

**Reference:** `/.windsurf/workflows/form-implementation.md`

### 8.4 Export Components

**File**: `src/features/{name}/components/index.ts`

**Tasks:**
1. ✅ Export all component files
2. ✅ Use named exports

---

## Step 9: Add Translation Keys

### Location
`packages/admin/src/assets/locales/{lang}/translation.json`

### English Keys (en)
**Tasks:**
1. ✅ Add menu label
2. ✅ Add breadcrumb label
3. ✅ Add feature labels (list, add, edit)
4. ✅ Add validation messages
5. ✅ Add success messages

### Lao Keys (lo)
**Tasks:**
1. ✅ Translate all English keys to Lao
2. ✅ Match structure exactly

### Required Keys Structure
```
{names}:
  - list, add, edit, filters
  - validation: { nameRequired, nameMaxLength, ... }
  - createSuccess, updateSuccess, deleteSuccess
menu:
  - {names}
breadcrumbs:
  - {names}, {name}Detail
```

### Reference
**See:** `/.windsurf/skills/feature/examples/complete-feature-structure.md` → Translation Files

---

## Step 10: Add to Navigation Menu

### Tasks
1. ✅ Add menu item to navigation config
2. ✅ Add icon import
3. ✅ Add path to feature route
4. ✅ Add permission check (optional)
5. ✅ Test menu visibility

### Reference
**See:** `/.windsurf/skills/feature/examples/menu-and-routing.md`

---

## Step 11: Export Public API

### Location
`src/features/{name}/index.ts`

### Tasks
1. ✅ Export components
2. ✅ Export hooks
3. ✅ Export types
4. ✅ Use named exports only

---

## Verification Checklist

### Structure
- [ ] Feature folder created with 6 subfolders
- [ ] All barrel exports (index.ts) in place
- [ ] Route file created
- [ ] Public API exported at feature root

### Files
- [ ] Constants file with endpoints and query keys
- [ ] Service file with all CRUD methods + signal
- [ ] Schema file as function accepting TFunction
- [ ] Types file with interfaces
- [ ] Hook files for queries and mutations
- [ ] Component files (List, Filters, Form)

### Implementation
- [ ] Query hooks use hierarchical query keys
- [ ] All service methods accept signal parameter
- [ ] Mutations invalidate correct queries
- [ ] NO try-catch in mutationFn
- [ ] Route has search params schema with z.undefined()
- [ ] Breadcrumb added to route

### Components
- [ ] List uses useMemo for columns
- [ ] List has dropdown row actions
- [ ] Filters sync with URL (if applicable)
- [ ] Form uses form.AppForm wrapper
- [ ] Form uses form.SubmitButton

### Translations
- [ ] English translations added
- [ ] Lao translations added
- [ ] Menu and breadcrumb keys added
- [ ] Validation messages use translation keys

### Integration
- [ ] Menu item added to navigation
- [ ] Feature accessible via menu
- [ ] Breadcrumbs display correctly

---

## Related Documentation

**Architecture & Structure:**
- `/.windsurf/skills/feature/SKILL.md` - Feature architecture concepts
- `/.windsurf/skills/feature/examples/` - Complete examples

**Implementation Workflows:**
- `/.windsurf/workflows/form-implementation.md` - Form implementation
- `/.windsurf/workflows/list-table-implementation.md` - List/table implementation

**Reference Implementation:**
- `src/features/positions/` - Complete working example

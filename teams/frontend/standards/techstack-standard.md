---
trigger: always_on
---

# Professional Project Standards (2026)


## Tech Stack (Verified Jan 2026)

- **Vite** 6.0.7 | **MUI** v7 | **Axios** 1.13.2
- **TanStack**: Router v1.136.8, Query v5.90.10, Form v1

- **Validation**: Zod 3.24.1 | **i18n**: i18next v24.2.2 | **Date**: Day.js 1.11.13

## Project Structure & Architecture
- **Feature-First**: All business logic must reside in `src/features/{feature-name}/`. 
- **Core Separation**: API, Router, and Theme config must remain in `src/core/`.
- **Routing**: Strictly use **File-Based Routing** in `src/routes/`. Use `__root.tsx` for global layouts.
- **Shared**: Only generic, non-business, re-use components (Buttons, Inputs, Dialogs) belong in `src/shared/`.

## Coding Constraints
- **Data Fetching**: Use `useQuery` or `useMutation`. Never use `useEffect` for fetching.
- **Form Handling**: Use `@tanstack/react-form` with `Zod` schemas. Wrap MUI components in a controller-like shared component in `src/shared/components/form/`. Create custome hook form composite
- **API**: Use the `axiosInstance` from `src/core/api/`. Pass `signal` to every request for cancellation.
- **Imports**: Use absolute paths (e.g., `@/features/auth/...`) if configured in Vite, otherwise relative within the same feature.

## Key Benefits Achieved:
# Single Responsibility:
- Each component has one clear purpose
- Easy to understand and modify
- Isolated bug fixes and feature additions
# Reusability:
- Components can be used in other parts of the app
- Hooks can be shared across booking-related features
- Clean interfaces for easy integration
# Maintainability:
- Changes are isolated to specific components
- New features can be added without affecting others
- Clear separation of concerns
# Testability:
- Small, focused components are easier to test
- Hooks can be tested independently
- Clear props interfaces for mocking
## Performance:
- Smaller re-render surfaces
- Components can be optimized individually
- Better React.memo opportunities
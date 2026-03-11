# Complete Feature Structure Example

This example shows the complete file structure for a **Product** feature with all preset files.

## Directory Structure

```
src/features/product/
├── components/
│   ├── ProductList.tsx
│   ├── ProductFilters.tsx
│   ├── ProductForm.tsx
│   └── index.ts
├── hooks/
│   ├── useProducts.ts
│   ├── useProduct.ts
│   ├── useCreateProduct.ts
│   ├── useUpdateProduct.ts
│   ├── useDeleteProduct.ts
│   └── index.ts
├── schemas/
│   └── product.schema.ts
├── services/
│   └── product.service.ts
├── types/
│   └── index.ts
├── constants/
│   └── product.constants.ts
└── index.ts

src/routes/_layout/products/
└── index.tsx
```

---

## File: `constants/product.constants.ts`

```typescript
// API Endpoints
export const PRODUCT_ENDPOINTS = {
  LIST: '/products',
  DETAIL: '/products/:id',
  CREATE: '/products',
  UPDATE: '/products/:id',
  DELETE: '/products/:id',
} as const;

// Query Keys (Hierarchical)
export const PRODUCT_QUERY_KEYS = {
  all: ['products'] as const,
  lists: () => [...PRODUCT_QUERY_KEYS.all, 'list'] as const,
  list: (params: any) => [...PRODUCT_QUERY_KEYS.lists(), params] as const,
  details: () => [...PRODUCT_QUERY_KEYS.all, 'detail'] as const,
  detail: (id: number) => [...PRODUCT_QUERY_KEYS.details(), id] as const,
} as const;

// Enums (Optional)
export const PRODUCT_STATUS = {
  ACTIVE: 'active',
  INACTIVE: 'inactive',
  DRAFT: 'draft',
} as const;

export type ProductStatus = typeof PRODUCT_STATUS[keyof typeof PRODUCT_STATUS];
```

**Purpose:**
- Centralize all static configuration
- Provide type-safe constants
- Define hierarchical query keys for cache management

---

## File: `types/index.ts`

```typescript
// Main entity interface
export interface Product {
  id: number;
  name: string;
  description: string | null;
  price: number;
  category_id: number;
  status: 'active' | 'inactive' | 'draft';
  created_at: string;
  updated_at: string;
}

// API response types
export interface ProductsResponse {
  products: Product[];
  total: number;
}

// Re-export form data type from schema
export type { ProductFormData } from '../schemas/product.schema';
```

**Purpose:**
- Define TypeScript interfaces for data structures
- Type API responses
- Re-export schema-generated types

---

## File: `schemas/product.schema.ts`

```typescript
import { z } from 'zod';
import type { TFunction } from 'i18next';

// Schema as function accepting TFunction for i18n
export const getProductSchema = (t: TFunction) => z.object({
  name: z
    .string()
    .min(1, t('products.validation.nameRequired'))
    .max(100, t('products.validation.nameMaxLength')),
  
  description: z
    .string()
    .max(500, t('products.validation.descriptionMaxLength'))
    .optional(),
  
  price: z
    .number()
    .min(0, t('products.validation.priceMin'))
    .max(999999, t('products.validation.priceMax')),
  
  category_id: z
    .number()
    .min(1, t('products.validation.categoryRequired')),
  
  status: z.enum(['active', 'inactive', 'draft']),
});

// Infer type from schema
export type ProductFormData = z.infer<ReturnType<typeof getProductSchema>>;
```

**Purpose:**
- Define validation rules with Zod
- Support internationalized error messages
- Generate TypeScript types automatically

---

## File: `services/product.service.ts`

```typescript
import axiosInstance from '@/core/api/axios.config';
import { PRODUCT_ENDPOINTS } from '../constants/product.constants';
import type { Product, ProductFormData, ProductsResponse } from '../types';

export const productService = {
  // GET /products - List with filters
  async getAll(params: any, signal?: AbortSignal): Promise<ProductsResponse> {
    const response = await axiosInstance.get<ProductsResponse>(
      PRODUCT_ENDPOINTS.LIST,
      { params, signal }
    );
    return response.data;
  },

  // GET /products/:id - Single item
  async getById(id: number, signal?: AbortSignal): Promise<{ product: Product }> {
    const response = await axiosInstance.get<{ product: Product }>(
      PRODUCT_ENDPOINTS.DETAIL.replace(':id', String(id)),
      { signal }
    );
    return response.data;
  },

  // POST /products - Create
  async create(data: ProductFormData, signal?: AbortSignal): Promise<{ product: Product }> {
    const response = await axiosInstance.post<{ product: Product }>(
      PRODUCT_ENDPOINTS.CREATE,
      data,
      { signal }
    );
    return response.data;
  },

  // PUT /products/:id - Update
  async update(
    id: number,
    data: ProductFormData,
    signal?: AbortSignal
  ): Promise<{ product: Product }> {
    const response = await axiosInstance.put<{ product: Product }>(
      PRODUCT_ENDPOINTS.UPDATE.replace(':id', String(id)),
      data,
      { signal }
    );
    return response.data;
  },

  // DELETE /products/:id - Delete
  async delete(id: number, signal?: AbortSignal): Promise<void> {
    await axiosInstance.delete(
      PRODUCT_ENDPOINTS.DELETE.replace(':id', String(id)),
      { signal }
    );
  },
};
```

**Purpose:**
- Handle all HTTP communication
- Accept AbortSignal for request cancellation
- Return typed responses
- No state management or UI logic

---

## File: `hooks/useProducts.ts`

```typescript
import { useQuery } from '@tanstack/react-query';
import { productService } from '../services/product.service';
import { PRODUCT_QUERY_KEYS } from '../constants/product.constants';
import type { ProductsSearch } from '@/routes/_layout/products';

export function useProducts(params: ProductsSearch) {
  return useQuery({
    queryKey: PRODUCT_QUERY_KEYS.list(params),
    queryFn: ({ signal }) => productService.getAll(params, signal),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}
```

**Purpose:**
- Wrap TanStack Query for list fetching
- Use hierarchical query keys
- Pass signal for automatic cancellation
- Configure cache behavior

---

## File: `hooks/useProduct.ts`

```typescript
import { useQuery } from '@tanstack/react-query';
import { productService } from '../services/product.service';
import { PRODUCT_QUERY_KEYS } from '../constants/product.constants';

export function useProduct(id: number, options?: { enabled?: boolean }) {
  return useQuery({
    queryKey: PRODUCT_QUERY_KEYS.detail(id),
    queryFn: ({ signal }) => productService.getById(id, signal),
    enabled: id > 0 && (options?.enabled ?? true),
    staleTime: 5 * 60 * 1000,
    ...options,
  });
}
```

**Purpose:**
- Fetch single item by ID
- Conditional fetching with enabled flag
- Prevent invalid queries (id <= 0)

---

## File: `hooks/useCreateProduct.ts`

```typescript
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { productService } from '../services/product.service';
import { PRODUCT_QUERY_KEYS } from '../constants/product.constants';
import type { ProductFormData } from '../types';

export function useCreateProduct() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: ProductFormData) => productService.create(data),
    onSuccess: () => {
      // Invalidate list queries to trigger refetch
      queryClient.invalidateQueries({ queryKey: PRODUCT_QUERY_KEYS.lists() });
    },
  });
}
```

**Purpose:**
- Handle create mutations
- Automatic cache invalidation
- No try-catch (errors handled by onError callback)

---

## File: `hooks/useUpdateProduct.ts`

```typescript
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { productService } from '../services/product.service';
import { PRODUCT_QUERY_KEYS } from '../constants/product.constants';
import type { ProductFormData } from '../types';

export function useUpdateProduct(id: number) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: ProductFormData) => productService.update(id, data),
    onSuccess: () => {
      // Invalidate all product queries (lists and details)
      queryClient.invalidateQueries({ queryKey: PRODUCT_QUERY_KEYS.all });
    },
  });
}
```

**Purpose:**
- Handle update mutations
- Accept ID as hook parameter
- Invalidate broader cache scope

---

## File: `hooks/useDeleteProduct.ts`

```typescript
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { productService } from '../services/product.service';
import { PRODUCT_QUERY_KEYS } from '../constants/product.constants';

export function useDeleteProduct() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => productService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: PRODUCT_QUERY_KEYS.lists() });
    },
  });
}
```

**Purpose:**
- Handle delete mutations
- Invalidate list queries after deletion

---

## File: `hooks/index.ts`

```typescript
export * from './useProducts';
export * from './useProduct';
export * from './useCreateProduct';
export * from './useUpdateProduct';
export * from './useDeleteProduct';
```

**Purpose:**
- Barrel export for clean imports
- Single import point for all hooks

---

## File: `components/index.ts`

```typescript
export * from './ProductList';
export * from './ProductFilters';
export * from './ProductForm';
```

**Purpose:**
- Barrel export for components
- Clean import path

---

## File: `index.ts` (Feature Root)

```typescript
// Public API - Export only what other features need
export * from './components';
export * from './hooks';
export * from './types';
```

**Purpose:**
- Define feature's public API
- Control what gets exposed
- Enable clean cross-feature imports (if needed)

---

## File: `routes/_layout/products/index.tsx`

```typescript
import { createFileRoute } from '@tanstack/react-router';
import { fallback, zodValidator } from '@tanstack/zod-adapter';
import { z } from 'zod';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Box, Button } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import { ProductList, ProductFilters, ProductForm } from '@/features/product';

// Search params schema
const productsSearchSchema = z.object({
  page: fallback(z.number(), 1).default(1),
  page_size: fallback(z.number(), 10).default(10),
  search: fallback(z.string(), z.undefined()).optional(),
  category_id: fallback(z.number(), z.undefined()).optional(),
  status: fallback(z.enum(['active', 'inactive', 'draft']), z.undefined()).optional(),
});

export type ProductsSearch = z.infer<typeof productsSearchSchema>;

// Route definition
export const Route = createFileRoute('/_layout/products/')({
  validateSearch: zodValidator(productsSearchSchema),
  component: ProductsPage,
});

// Page component
function ProductsPage() {
  const { t } = useTranslation();
  const searchParams = Route.useSearch();
  const navigate = Route.useNavigate();
  
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [editingId, setEditingId] = useState<number | null>(null);

  const handleAddClick = () => {
    setEditingId(null);
    setIsFormOpen(true);
  };

  const handleEditClick = (productId: number) => {
    setEditingId(productId);
    setIsFormOpen(true);
  };

  const handleFormClose = () => {
    setIsFormOpen(false);
    setEditingId(null);
  };

  const handlePageChange = (page: number) => {
    navigate({ search: (prev) => ({ ...prev, page }) });
  };

  const handlePageSizeChange = (pageSize: number) => {
    navigate({ search: (prev) => ({ ...prev, page: 1, page_size: pageSize }) });
  };

  return (
    <Box sx={{ p: { xs: 2, md: 4 } }}>
      <Box sx={{ maxWidth: 1440, mx: 'auto' }}>
        {/* Header */}
        <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
          <Box component="h1" sx={{ fontSize: { xs: '1.5rem', md: '2rem' }, fontWeight: 600 }}>
            {t('products.list')}
          </Box>
          <Button
            variant="contained"
            startIcon={<AddIcon />}
            onClick={handleAddClick}
            sx={{ minHeight: 48 }}
          >
            {t('products.add')}
          </Button>
        </Box>

        {/* Filters */}
        <Box sx={{ mb: 3 }}>
          <ProductFilters />
        </Box>

        {/* List */}
        <ProductList
          searchParams={searchParams}
          onPageChange={handlePageChange}
          onPageSizeChange={handlePageSizeChange}
          onEditClick={handleEditClick}
        />

        {/* Form Dialog */}
        <ProductForm
          open={isFormOpen}
          productId={editingId}
          onClose={handleFormClose}
        />
      </Box>
    </Box>
  );
}
```

**Purpose:**
- Define route with search params validation
- Orchestrate all feature components
- Manage page-level state
- Handle navigation

---

## Translation Files

### `locales/en/translation.json`

```json
{
  "products": {
    "list": "Products",
    "add": "Add Product",
    "edit": "Edit Product",
    "filters": "Filters",
    "validation": {
      "nameRequired": "Product name is required",
      "nameMaxLength": "Name must be at most 100 characters",
      "descriptionMaxLength": "Description must be at most 500 characters",
      "priceMin": "Price must be at least 0",
      "priceMax": "Price must be at most 999,999",
      "categoryRequired": "Category is required"
    },
    "createSuccess": "Product created successfully",
    "updateSuccess": "Product updated successfully",
    "deleteSuccess": "Product deleted successfully"
  }
}
```

### `locales/lo/translation.json`

```json
{
  "products": {
    "list": "ສິນຄ້າ",
    "add": "ເພີ່ມສິນຄ້າ",
    "edit": "ແກ້ໄຂສິນຄ້າ",
    "filters": "ຕົວກອງ",
    "validation": {
      "nameRequired": "ກະລຸນາໃສ່ຊື່ສິນຄ້າ",
      "nameMaxLength": "ຊື່ຕ້ອງບໍ່ເກີນ 100 ຕົວອັກສອນ",
      "descriptionMaxLength": "ລາຍລະອຽດຕ້ອງບໍ່ເກີນ 500 ຕົວອັກສອນ",
      "priceMin": "ລາຄາຕ້ອງຫຼາຍກວ່າ 0",
      "priceMax": "ລາຄາຕ້ອງບໍ່ເກີນ 999,999",
      "categoryRequired": "ກະລຸນາເລືອກປະເພດ"
    },
    "createSuccess": "ສ້າງສິນຄ້າສຳເລັດ",
    "updateSuccess": "ແກ້ໄຂສິນຄ້າສຳເລັດ",
    "deleteSuccess": "ລຶບສິນຄ້າສຳເລັດ"
  }
}
```

---

## Key Principles Demonstrated

### 1. **Feature Encapsulation**
- All product-related code in one folder
- Self-contained and independent
- Clear boundaries

### 2. **Layer Separation**
- Components → Hooks → Services → API
- Each layer has specific responsibility
- No layer skipping

### 3. **Type Safety**
- Schema generates types
- Service methods are typed
- Query/mutation hooks are typed

### 4. **Internationalization**
- Schema accepts TFunction
- All UI text uses translation keys
- Both English and Lao supported

### 5. **Cache Management**
- Hierarchical query keys
- Proper invalidation strategy
- Optimized refetching

### 6. **Clean Imports**
- Barrel exports in each folder
- Public API at feature root
- Absolute imports from feature

---

## Component Implementation

For actual component implementation (List, Filters, Form), refer to:
- **List/Table**: `/.windsurf/skills/list-table/SKILL.md`
- **Form**: `/.windsurf/skills/form/SKILL.md`
- **Workflows**: `/.windsurf/workflows/`

This example focuses on **structure and file organization**, not implementation details.

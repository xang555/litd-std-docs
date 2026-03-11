# Menu Integration & Routing Example

Complete example showing how to integrate a feature into the application's navigation menu and routing system.

## Scenario: Adding Products Feature

We'll add a Products feature to the application with:
- Sidebar menu item
- List page route
- Detail page route
- Breadcrumb navigation
- Permission-based access

---

## Step 1: Create Route Files

### List Page Route
**File:** `src/routes/_layout/products/index.tsx`

```typescript
import { createFileRoute } from '@tanstack/react-router';
import { fallback, zodValidator } from '@tanstack/zod-adapter';
import { z } from 'zod';
import { ProductsPage } from './ProductsPage';

const productsSearchSchema = z.object({
  page: fallback(z.number(), 1).default(1),
  page_size: fallback(z.number(), 10).default(10),
  search: fallback(z.string(), z.undefined()).optional(),
  category_id: fallback(z.number(), z.undefined()).optional(),
});

export type ProductsSearch = z.infer<typeof productsSearchSchema>;

export const Route = createFileRoute('/_layout/products/')({
  validateSearch: zodValidator(productsSearchSchema),
  beforeLoad: () => ({
    breadcrumb: 'breadcrumbs.products',
  }),
  component: ProductsPage,
});
```

**URL:** `/products`

### Detail Page Route (Optional)
**File:** `src/routes/_layout/products/$id.tsx`

```typescript
import { createFileRoute } from '@tanstack/react-router';
import { ProductDetailPage } from './ProductDetailPage';

export const Route = createFileRoute('/_layout/products/$id')({
  beforeLoad: ({ params }) => ({
    breadcrumb: 'breadcrumbs.productDetail',
  }),
  component: ProductDetailPage,
});
```

**URL:** `/products/:id`

---

## Step 2: Add Menu Item

### Option A: Direct Menu Configuration
**File:** `src/core/layout/navigation/menuConfig.ts` (or similar)

```typescript
import InventoryIcon from '@mui/icons-material/Inventory';

export const menuItems = [
  // ... other menu items
  {
    id: 'products',
    label: 'menu.products',
    icon: InventoryIcon,
    path: '/products',
    permission: 'products.view',
  },
];
```

### Option B: Grouped Menu Items
```typescript
export const menuItems = [
  // ... other menu items
  {
    id: 'inventory',
    label: 'menu.inventory',
    icon: InventoryIcon,
    children: [
      {
        id: 'products',
        label: 'menu.products',
        path: '/products',
        permission: 'products.view',
      },
      {
        id: 'categories',
        label: 'menu.categories',
        path: '/categories',
        permission: 'categories.view',
      },
    ],
  },
];
```

### Option C: Sidebar Component Integration
**File:** `src/core/layout/Sidebar.tsx`

```typescript
import { useTranslation } from 'react-i18next';
import { List, ListItem, ListItemButton, ListItemIcon, ListItemText } from '@mui/material';
import InventoryIcon from '@mui/icons-material/Inventory';
import { Link } from '@tanstack/react-router';

export function Sidebar() {
  const { t } = useTranslation();

  return (
    <List>
      {/* ... other menu items */}
      
      <ListItem disablePadding>
        <ListItemButton component={Link} to="/products">
          <ListItemIcon>
            <InventoryIcon />
          </ListItemIcon>
          <ListItemText primary={t('menu.products')} />
        </ListItemButton>
      </ListItem>
    </List>
  );
}
```

---

## Step 3: Add Translation Keys

### English
**File:** `packages/admin/public/locales/en/translation.json`

```json
{
  "menu": {
    "products": "Products",
    "inventory": "Inventory Management",
    "categories": "Categories"
  },
  "breadcrumbs": {
    "home": "Home",
    "products": "Products",
    "productDetail": "Product Details"
  },
  "products": {
    "list": "Products",
    "add": "Add Product",
    "edit": "Edit Product"
  }
}
```

### Lao
**File:** `packages/admin/public/locales/lo/translation.json`

```json
{
  "menu": {
    "products": "ສິນຄ້າ",
    "inventory": "ການຈັດການສິນຄ້າ",
    "categories": "ປະເພດ"
  },
  "breadcrumbs": {
    "home": "ໜ້າຫຼັກ",
    "products": "ສິນຄ້າ",
    "productDetail": "ລາຍລະອຽດສິນຄ້າ"
  },
  "products": {
    "list": "ສິນຄ້າ",
    "add": "ເພີ່ມສິນຄ້າ",
    "edit": "ແກ້ໄຂສິນຄ້າ"
  }
}
```

---

## Step 4: Navigation Actions in Components

### From List to Detail
**In:** `ProductList.tsx`

```typescript
import { useNavigate } from '@tanstack/react-router';

export function ProductList({ searchParams, onEditClick }: Props) {
  const navigate = useNavigate();

  const handleViewDetails = (productId: number) => {
    navigate({
      to: '/products/$id',
      params: { id: productId.toString() },
    });
  };

  // In row actions dropdown
  <MenuItem onClick={() => handleViewDetails(row.original.id)}>
    <ListItemIcon><VisibilityIcon fontSize="small" /></ListItemIcon>
    <ListItemText>{t('common.view')}</ListItemText>
  </MenuItem>
}
```

### From Detail Back to List
**In:** `ProductDetailPage.tsx`

```typescript
import { useNavigate } from '@tanstack/react-router';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

export function ProductDetailPage() {
  const navigate = useNavigate();

  const handleBackToList = () => {
    navigate({ to: '/products' });
  };

  return (
    <Box>
      <Button
        startIcon={<ArrowBackIcon />}
        onClick={handleBackToList}
      >
        {t('common.back')}
      </Button>
      {/* Detail content */}
    </Box>
  );
}
```

### After Create Success
**In:** `ProductForm.tsx`

```typescript
const createMutation = useCreateProduct();

const handleSubmit = (formData: ProductFormData) => {
  createMutation.mutate(formData, {
    onSuccess: (response) => {
      showSuccess(t('products.createSuccess'));
      
      // Option A: Close dialog and stay on list
      onClose();
      
      // Option B: Navigate to new product detail
      navigate({
        to: '/products/$id',
        params: { id: response.product.id.toString() },
      });
    },
    onError: (error) => {
      showError(translateError(error));
    },
  });
};
```

### Preserve Search Params
**In:** `ProductList.tsx`

```typescript
const handleEdit = (productId: number) => {
  // Keep current filters when navigating
  navigate({
    to: '/products/$id/edit',
    params: { id: productId.toString() },
    search: (prev) => prev, // Preserve search params
  });
};
```

---

## Step 5: Permission-Based Access

### Menu Item with Permission Check
```typescript
import { usePermissions } from '@/core/auth/hooks/usePermissions';

export function Sidebar() {
  const { t } = useTranslation();
  const { hasPermission } = usePermissions();

  const menuItems = [
    {
      label: t('menu.products'),
      icon: <InventoryIcon />,
      path: '/products',
      permission: 'products.view',
      visible: hasPermission('products.view'),
    },
    {
      label: t('menu.settings'),
      icon: <SettingsIcon />,
      permission: 'settings.view',
      visible: hasPermission('settings.view'),
      children: [
        {
          label: t('menu.categories'),
          path: '/settings/categories',
          permission: 'categories.manage',
          visible: hasPermission('categories.manage'),
        },
      ],
    },
  ];

  return (
    <List>
      {menuItems
        .filter(item => item.visible)
        .map(item => (
          <MenuItem key={item.path} item={item} />
        ))}
    </List>
  );
}
```

### Route-Level Permission Check
```typescript
export const Route = createFileRoute('/_layout/products/')({
  validateSearch: zodValidator(productsSearchSchema),
  beforeLoad: ({ context }) => {
    // Check permission before loading route
    if (!context.auth.hasPermission('products.view')) {
      throw redirect({ to: '/unauthorized' });
    }
    
    return {
      breadcrumb: 'breadcrumbs.products',
    };
  },
  component: ProductsPage,
});
```

---

## Step 6: Breadcrumb Integration

### Breadcrumb Component
**File:** `src/core/layout/Breadcrumbs.tsx`

```typescript
import { useMatches } from '@tanstack/react-router';
import { Breadcrumbs as MuiBreadcrumbs, Link, Typography } from '@mui/material';
import { useTranslation } from 'react-i18next';
import HomeIcon from '@mui/icons-material/Home';

export function Breadcrumbs() {
  const { t } = useTranslation();
  const matches = useMatches();

  const breadcrumbs = matches
    .filter(match => match.context?.breadcrumb)
    .map(match => ({
      label: t(match.context.breadcrumb),
      path: match.pathname,
    }));

  return (
    <MuiBreadcrumbs>
      <Link href="/" sx={{ display: 'flex', alignItems: 'center' }}>
        <HomeIcon sx={{ mr: 0.5 }} fontSize="small" />
        {t('breadcrumbs.home')}
      </Link>
      
      {breadcrumbs.map((crumb, index) => {
        const isLast = index === breadcrumbs.length - 1;
        
        return isLast ? (
          <Typography key={crumb.path} color="text.primary">
            {crumb.label}
          </Typography>
        ) : (
          <Link key={crumb.path} href={crumb.path}>
            {crumb.label}
          </Link>
        );
      })}
    </MuiBreadcrumbs>
  );
}
```

---

## Complete Route Structure Example

```
src/routes/
├── __root.tsx                      # Root layout
├── _layout/                        # Main app layout
│   ├── index.tsx                   # Dashboard (/)
│   ├── products/
│   │   ├── index.tsx               # /products (list)
│   │   ├── $id.tsx                 # /products/:id (detail)
│   │   └── $id.edit.tsx            # /products/:id/edit (edit)
│   ├── categories/
│   │   └── index.tsx               # /categories
│   ├── settings/
│   │   ├── index.tsx               # /settings
│   │   ├── users/
│   │   │   └── index.tsx           # /settings/users
│   │   └── roles/
│   │       └── index.tsx           # /settings/roles
│   └── profile/
│       └── index.tsx                # /profile
├── login.tsx                       # /login (no layout)
└── unauthorized.tsx                # /unauthorized
```

**Generated URLs:**
- `/` - Dashboard
- `/products` - Products list
- `/products/123` - Product detail
- `/products/123/edit` - Edit product
- `/categories` - Categories list
- `/settings` - Settings page
- `/settings/users` - User management
- `/login` - Login page

---

## Common Navigation Patterns Summary

### 1. List → Detail
```typescript
navigate({ to: '/products/$id', params: { id: '123' } });
```

### 2. Detail → List
```typescript
navigate({ to: '/products' });
```

### 3. List → Edit (preserve filters)
```typescript
navigate({
  to: '/products/$id/edit',
  params: { id: '123' },
  search: (prev) => prev,
});
```

### 4. After Create → Detail
```typescript
navigate({ to: '/products/$id', params: { id: newId.toString() } });
```

### 5. After Update → Stay
```typescript
// No navigation, cache invalidation refreshes data automatically
onClose();
```

### 6. After Delete → Stay
```typescript
// No navigation, cache invalidation refreshes list
```

---

## Checklist

When adding a new feature to navigation:

### Route Setup
- [ ] Create route file in `src/routes/_layout/{feature}/index.tsx`
- [ ] Define search params schema with Zod
- [ ] Add breadcrumb in `beforeLoad`
- [ ] Export route component

### Menu Integration
- [ ] Add menu item to navigation config
- [ ] Add icon import
- [ ] Add permission check (if needed)
- [ ] Test menu visibility

### Translations
- [ ] Add menu label translations (en, lo)
- [ ] Add breadcrumb translations (en, lo)
- [ ] Add feature-specific translations

### Navigation Actions
- [ ] Implement list → detail navigation
- [ ] Implement detail → list back button
- [ ] Handle post-create navigation
- [ ] Preserve search params when needed

### Permissions
- [ ] Define feature permissions
- [ ] Add permission checks to menu
- [ ] Add permission checks to routes (optional)
- [ ] Test unauthorized access

---

## Related Documentation

- **Feature Architecture**: `../SKILL.md`
- **Complete Feature Example**: `./complete-feature-structure.md`
- **TanStack Router Docs**: https://tanstack.com/router

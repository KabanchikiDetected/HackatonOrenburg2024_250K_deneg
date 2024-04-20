# Docs

## Tree

```
├── public  
│   ├── fonts  
│   └── images  
├── src  
│   ├── app  
│   │   ├── Router  
│   │   ├── contexts  
│   │   ├── styles  
│   │   ├── index.scss  
│   │   └── index.tsx  
│   ├── features  
│   │   ├── components  
│   │   ├── hooks  
│   │   └── index.tsx  
│   ├── main.tsx  
│   ├── pages  
│   ├── shared  
│   └── widgets  
├── tsconfig.json  
├── tsconfig.node.json  
└── vite.config.ts  
```

## public/

This folder contains image and font directories

## src/

This folder contains source code

### src/app/

This folder contains app settings

### src/features/

This folder contains hooks and components directories

### src/pages/

This directory contains subdirectories with components for pages

### src/shared/
This directory contains tools that will remain unchanged

### src/main.tsx

Main file where mounted App

## Components

### Loader

This component contains loader spinner

**using:**
```javascript
<Loader children={<Component />} />
```

or 

```javascript
<Loader>
  <Component />
</Loader>
```

`Component` – a component that needs to be loaded, and a spinner will be shown during the loading process.

### Card

**using:**
```javascript
<Card data={DATA} onSwipe={eventHandle} />
```

`DATA` – a variable contains data for card

interface (for DATA):
```typescript
interface IData {
  title: string;
  tags: string[];
  img: string;
  about: string;
}
```

`onSwipe` – a function that must execute when a card is swiped left or right
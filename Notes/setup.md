# Setting Up The Application Yourself
## Folders
Create separate ``frontend`` and ``backend`` folders.

## React frontend
### Creating Next.js project
In terminal run: ``npx create-next-app@latest .``


Current Setup:

1. √ Would you like to use TypeScript? ... No / **Yes**
2. √ Would you like to use ESLint? ... **No** / Yes
3. √ Would you like to use Tailwind CSS? ... No / **Yes**
4. √ Would you like your code inside a `src/` directory? ... No / **Yes**
5. √ Would you like to use App Router? (recommended) ... No / **Yes**
6. √ Would you like to use Turbopack for `next dev`? ... **No** / Yes
7. ? Would you like to customize the import alias (`@/*` by default)? » **No** / Yes


### Creating app folder
Path: ``/AppName/frontend/app/``

Inside the app directory, create ``layout.tsx`` and ``page.tsx``.
These files define the root route (``/``) providing consistent structure across the route.

Each folder inside app (``about/``, ``contact/`` etc.) automatically becomes a route as soon as a page.tsx file
is created inside it.

#### layout.tsx
1. Shared layout for a specific route or the entire app.
2. Used for navigation bars, footers, or other global styles.

#### page.tsx
1. Defines the content for a specific route


### Extra stuff
1. Check the ``next.config.ts`` file for comments.






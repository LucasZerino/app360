import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;
  const isPublicPath = pathname === '/' || 
                      pathname === '/auth/login' || 
                      pathname === '/auth/register' || 
                      pathname === '/auth/forgot-password';

  const token = request.cookies.get('auth-token')?.value;

  // Se não estiver autenticado e tentar acessar rota protegida
  if (!token && !isPublicPath) {
    return NextResponse.redirect(new URL('/auth/login', request.url));
  }

  // Se estiver autenticado e tentar acessar rotas de auth
  if (token && pathname.startsWith('/auth')) {
    return NextResponse.redirect(new URL('/', request.url));
  }

  return NextResponse.next();
}

// Configurar quais rotas o middleware deve interceptar
export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     * - public (public files)
     * - assets (assets files)
     */
    '/((?!api|_next/static|_next/image|favicon.ico|public|assets).*)',
  ],
}; 
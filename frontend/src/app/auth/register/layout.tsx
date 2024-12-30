'use client';

// components
import AuthModernLayout from '@/layouts/auth/modern';

// ----------------------------------------------------------------------

type Props = {
  children: React.ReactNode;
};

export default function Layout({ children }: Props) {
  return <AuthModernLayout>{children}</AuthModernLayout>;
}
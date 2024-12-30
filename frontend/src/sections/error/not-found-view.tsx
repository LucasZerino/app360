'use client';

import { m } from 'framer-motion';
// @mui
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
// layouts
import CompactLayout from '@/layouts/compact';
// routes
import { RouterLink } from '@/routes/components';
// components
import { MotionContainer, varBounce } from '@/components/animate';
// assets
import { PageNotFoundIllustration } from '@/assets/illustrations';

// ----------------------------------------------------------------------

export default function NotFoundView() {
  return (
    <CompactLayout>
      <MotionContainer>
        <m.div variants={varBounce().in}>
          <Typography variant="h3" paragraph>
            Página não encontrada!
          </Typography>
        </m.div>

        <m.div variants={varBounce().in}>
          <Typography sx={{ color: 'text.secondary' }}>
            Desculpe, não conseguimos encontrar a página que você está procurando. 
            Talvez você tenha digitado incorretamente a URL? Verifique se a escrita está correta.
          </Typography>
        </m.div>

        <m.div variants={varBounce().in}>
          <PageNotFoundIllustration
            sx={{
              height: 260,
              my: { xs: 5, sm: 10 },
            }}
          />
        </m.div>

        <Button component={RouterLink} href="/" size="large" variant="contained">
          Ir para Início
        </Button>
      </MotionContainer>
    </CompactLayout>
  );
}

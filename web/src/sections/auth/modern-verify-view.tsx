'use client';

import * as Yup from 'yup';
import { useCallback } from 'react';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
// @mui
import LoadingButton from '@mui/lab/LoadingButton';
import Link from '@mui/material/Link';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
// routes
import { paths } from '@/routes/paths';
import { RouterLink } from '@/routes/components';
// components
import Iconify from '@/components/iconify';
import FormProvider, { RHFCode, RHFTextField } from '@/components/hook-form';
// assets
import { EmailInboxIcon } from '@/assets/icons';

// ----------------------------------------------------------------------

type FormValuesProps = {
  code: string;
  email: string;
};

export default function ModernVerifyView() {
  const VerifySchema = Yup.object().shape({
    code: Yup.string()
      .min(6, 'O código deve ter pelo menos 6 caracteres')
      .required('O código é obrigatório'),
    email: Yup.string()
      .required('E-mail é obrigatório')
      .email('O e-mail deve ser um endereço válido'),
  });

  const defaultValues = {
    code: '',
    email: '',
  };

  const methods = useForm({
    mode: 'onChange',
    resolver: yupResolver(VerifySchema),
    defaultValues,
  });

  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const onSubmit = useCallback(async (data: FormValuesProps) => {
    try {
      await new Promise((resolve) => setTimeout(resolve, 500));
      console.info('DATA', data);
    } catch (error) {
      console.error(error);
    }
  }, []);

  const renderForm = (
    <Stack spacing={3} alignItems="center">
      <RHFTextField
        name="email"
        label="E-mail"
        placeholder="exemplo@gmail.com"
        InputLabelProps={{ shrink: true }}
      />

      <RHFCode name="code" />

      <LoadingButton
        fullWidth
        size="large"
        type="submit"
        variant="contained"
        loading={isSubmitting}
      >
        Verificar
      </LoadingButton>

      <Typography variant="body2">
        {'Não recebeu o código? '}
        <Link
          variant="subtitle2"
          sx={{
            cursor: 'pointer',
          }}
        >
          Reenviar código
        </Link>
      </Typography>

      <Link
        component={RouterLink}
        href={paths.auth.login}
        color="inherit"
        variant="subtitle2"
        sx={{
          alignItems: 'center',
          display: 'inline-flex',
        }}
      >
        <Iconify icon="eva:arrow-ios-back-fill" width={16} />
        Voltar para o login
      </Link>
    </Stack>
  );

  const renderHead = (
    <>
      <EmailInboxIcon sx={{ height: 96 }} />

      <Stack spacing={1} sx={{ my: 5 }}>
        <Typography variant="h3">Por favor, verifique seu e-mail!</Typography>

        <Typography variant="body2" sx={{ color: 'text.secondary' }}>
          Enviamos um código de confirmação de 6 dígitos para acb@domain. 
          Por favor, insira o código abaixo para verificar seu e-mail.
        </Typography>
      </Stack>
    </>
  );

  return (
    <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
      {renderHead}

      {renderForm}
    </FormProvider>
  );
}

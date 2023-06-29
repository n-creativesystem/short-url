import { AsyncButton, CopyButton } from '@/components/Parts/Button';
import { Form, Input } from '@/components/Parts/Form';
import { Box } from '@mui/material';
import { FC, memo, useCallback } from 'react';
import { Controller, SubmitHandler, useForm } from 'react-hook-form';
import type { Data, Input as TInput } from './index.d';
type Props = {
  data?: Data;
  onClick: (input: TInput) => Promise<void>;
};

const initialValue = () => ({
  name: '',
  id: '',
  secret: '',
});

export const Presenter: FC<Props> = memo(({ data, onClick: onClickInput }) => {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<Data>({
    defaultValues: initialValue(),
    values: data,
  });
  const onSubmit: SubmitHandler<Data> = useCallback(
    (data) => {
      return onClickInput({ ...data });
    },
    [data]
  );
  const visibleCopy = data?.id !== '';

  const validationRules = {
    name: {
      required: 'アプリ名を入力してください',
      maxLength: { value: 255, message: '255文字以内で入力してください' },
    },
  };

  return (
    <Form
      noValidate
      onSubmit={handleSubmit(onSubmit)}
      spacing={4}
      sx={{ width: '1000px' }}
    >
      <Controller
        name="id"
        control={control}
        render={({ field }) => {
          return (
            <Box
              sx={{
                display: 'flex',
                alignItems: 'flex-end',
                columnGap: '8px',
              }}
            >
              <Input
                {...field}
                type="text"
                label="Client ID"
                error={errors.id !== undefined}
                helperText={errors.id?.message}
                fullWidth
                disabled
              />
              <CopyButton
                visible={visibleCopy}
                value={field.value}
                toolTipText="Client IDをコピー"
                size="small"
                variant="outlined"
              />
            </Box>
          );
        }}
      />
      <Controller
        name="secret"
        control={control}
        render={({ field }) => {
          return (
            <Box
              sx={{
                display: 'flex',
                alignItems: 'flex-end',
                columnGap: '8px',
              }}
            >
              <Input
                {...field}
                type="password"
                label="Client secret"
                error={errors.secret !== undefined}
                helperText={errors.secret?.message}
                fullWidth
                disabled
              />
              <CopyButton
                visible={visibleCopy}
                value={field.value}
                toolTipText="Client secretをコピー"
                size="small"
                variant="outlined"
              />
            </Box>
          );
        }}
      />
      <Controller
        name="name"
        control={control}
        rules={validationRules.name}
        render={({ field }) => {
          return (
            <>
              <Input
                {...field}
                type="text"
                label="アプリ名"
                error={errors.name !== undefined}
                helperText={errors.name?.message}
                required
              />
            </>
          );
        }}
      />
      <Box
        sx={{
          marginLeft: 'auto !important',
        }}
      >
        <AsyncButton variant="contained" type="submit">
          保存
        </AsyncButton>
      </Box>
    </Form>
  );
});

if (process.env.NODE_ENV !== 'production') {
  Presenter.displayName = 'OAuthAppContentPresenter';
}

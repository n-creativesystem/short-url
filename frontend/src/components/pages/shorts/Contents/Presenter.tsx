import { AsyncButton } from '@/components/Parts/Button';
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
  key: '',
  url: '',
  created_at: '',
  updated_at: '',
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

  const validationRules = {
    url: {
      required: 'urlを入力してください',
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
        name="key"
        control={control}
        render={({ field }) => {
          return (
            <Input {...field} type="text" label="key" fullWidth disabled />
          );
        }}
      />
      <Controller
        name="url"
        control={control}
        rules={validationRules.url}
        render={({ field }) => {
          return (
            <Input
              {...field}
              type="text"
              label="URL"
              error={errors.url !== undefined}
              helperText={errors.url?.message}
              fullWidth
              disabled
            />
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
  Presenter.displayName = 'ShortsContentPresenter';
}

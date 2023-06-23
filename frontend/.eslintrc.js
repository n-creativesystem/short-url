module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: [
    // @see https://github.com/vercel/next.js/blob/canary/packages/eslint-config-next/index.js
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:react-hooks/recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:prettier/recommended',
  ],
  plugins: ['@typescript-eslint', 'react-refresh'],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 'es2021',
    sourceType: 'module',
  },
  settings: {
    react: {
      version: 'detect',
    },
    'import/resolver': {
      node: {
        extensions: ['.js', '.jsx', '.json', '.ts', '.tsx'],
      },
      typescript: {},
    },
  },
  // @refs: https://gist.github.com/sin-tanaka/b18bf1b5b46bd685fee93bd26fb473b3
  rules: {
    // 関数の戻り値はtsの推論に任せる (exportする関数は必要)
    '@typescript-eslint/explicit-function-return-type': 'off',
    // anyを禁止 (必要なケースは行コメントでeslint-disableする)
    '@typescript-eslint/no-explicit-any': 'error',
    // ts-ignoreを許可する
    '@typescript-eslint/ban-ts-comment': 'off',
    // type Props = {} などを許可する ()
    '@typescript-eslint/ban-types': [
      'off',
      {
        types: {
          '{}': false,
        },
      },
    ],
    // 厳密等価演算子を強制
    eqeqeq: 2,
    // imgタグの利用を許可する
    '@next/next/no-img-element': 'off',
    'no-console': 'warn',
    // e.g. prop={'foo'} -> prop='foo'
    'react/jsx-curly-brace-presence': 'warn',
    // e.g. opened={true} -> opened
    'react/jsx-boolean-value': 'warn',
    // e.g. <Foo></Foo> -> <Foo />
    'react/self-closing-comp': [
      'warn',
      {
        component: true,
        html: true,
      },
    ],
    'react-refresh/only-export-components': 'warn',
  },
  overrides: [
    {
      files: ['*.ts', '*.tsx', '*.d.ts'],
    },
  ],
};

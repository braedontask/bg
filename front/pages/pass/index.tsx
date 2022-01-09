import type { NextPage } from 'next'
import Head from 'next/head'
import styles from './styles/Pass.module.css'
import React, { ChangeEvent } from 'react'
import { Button, Container, Form } from 'react-bootstrap'
import axios from 'axios'

interface IState {
  domain?: string,
  masterPassword?: string,
  version?: number,
  password?: string
}

const Pass: NextPage = () => {
  const [state, setState] = React.useState<IState | undefined>({
    domain: undefined,
    masterPassword: undefined,
    version: undefined,
    password: undefined
  })

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    setState({ ...state, [e.target.id]: e.target.value })
  }

  async function handleSubmit(e: React.SyntheticEvent) {
    if (state?.domain && state?.masterPassword && state?.version) {
      axios.post('/api/pass', {
        state
      }).then(function (res) {
        setState({ ...state, password: res.data })
        })
      e?.preventDefault()
      setState({ ...state, password: state?.masterPassword })
    }
  }

  return (
    <div>
      <Head>
        <title>Password Manager</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Container className={styles.container}>
        <h1>Password Manager</h1>
        <Form onSubmit={handleSubmit}>
          <Form.Group className={styles.formGroup} controlId="domain">
            <Form.Label>Domain</Form.Label>
            <Form.Control type="text" placeholder="amazon.com" onChange={handleChange}/>
          </Form.Group>
          <Form.Group className={styles.formGroup} controlId="masterPassword">
            <Form.Label>Master Password</Form.Label>
            <Form.Control type="password" placeholder="password"  onChange={handleChange}/>
          </Form.Group>
          <Form.Group className={styles.formGroup} controlId="version">
            <Form.Label>Version</Form.Label>
            <Form.Control type="number" placeholder="5" onChange={handleChange}/>
          </Form.Group>
          <Button type='submit'>Submit</Button>
        </Form>
        { state?.password }
      </Container>
    </div>
  )
}

export default Pass

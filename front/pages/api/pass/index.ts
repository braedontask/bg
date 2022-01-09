// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import axios from 'axios'

type Data = {
  password: string
}

export default function passwordHandler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  axios.post('http://localhost:8090/pass', {
    data: req.body
  }).then(function (resp) {
    res.status(200).json(resp.data)
  })
}

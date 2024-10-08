package service

import (
	"errors"
	dto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceDirectoryTree struct {
	directoryTreeRepository interfaces.IDirectoryRepository
	documentService         interfaces.IDocumentService
}

func NewServiceDirectory(r interfaces.IDirectoryRepository, documentService interfaces.IDocumentService) interfaces.IDirectoryService {
	return &serviceDirectoryTree{
		directoryTreeRepository: r,
		documentService:         documentService,
	}
}

func (s *serviceDirectoryTree) Get(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	return s.directoryTreeRepository.GetDirectory(req)
}

func (s *serviceDirectoryTree) Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error) {

	repoReq := &sq.CreateDirectoryTreeParams{
		Name:      pgtype.Text{String: req.Name, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	if req.ParentID > 0 {
		repoReq.ParentID = pgtype.Int8{Int64: req.ParentID, Valid: true}
	}
	if req.InstitutionID > 0 {
		repoReq.InstitutionID = pgtype.Int4{Int32: req.InstitutionID, Valid: true}
	}

	if req.AsesorID > 0 {
		repoReq.AsesorID = pgtype.Int4{Int32: req.AsesorID, Valid: true}
	}
	r, err := s.directoryTreeRepository.CreateDirectory(repoReq)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *serviceDirectoryTree) List() (*[]dto.GetDirectoryTreeRes, error) {
	repoRes, err := s.directoryTreeRepository.ListDirectory()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (s *serviceDirectoryTree) ListByParent(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	repoRes, err := s.directoryTreeRepository.ListDirectoryByParent(req)
	if err != nil {
		return nil, err
	}
	documents, err := s.documentService.ListByDirectory(dto.GetDirectoryTreeReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	directory, err := s.Get(req)
	if err != nil {
		return nil, err
	}

	return &dto.GetDirectoryTreeRes{
		Id:            directory.Id,
		Name:          directory.Name,
		Open:          true,
		ParentID:      directory.ParentID,
		InstitutionID: directory.InstitutionID,
		Directory:     repoRes, Document: documents,
	}, nil
}

func (s *serviceDirectoryTree) GetRoute(req *dto.GetDirectoryTreeReq) (*[]int64, *[]dto.GetDirectoryTreeRes, error) {
	directory, err := s.ListByParent(req)
	if err != nil {
		return nil, nil, err
	}
	repoRes, err := s.directoryTreeRepository.ListDirectoryHierarchy(req)
	if err != nil {
		return nil, nil, err
	}
	tree := []int64{directory.Id, directory.ParentID}
	var nodes []dto.GetDirectoryTreeRes

	nodes = append(nodes,
		dto.GetDirectoryTreeRes{
			Id:            directory.Id,
			Name:          directory.Name,
			ParentID:      directory.Id,
			InstitutionID: directory.InstitutionID,
		},
	)

	node := directory.ParentID
	for node != 0 {
		for _, directory_tree := range *repoRes {
			if directory_tree.Id == node && directory_tree.ParentID != 0 {
				nodes = append(nodes, dto.GetDirectoryTreeRes{
					Id:       directory_tree.Id,
					Name:     directory_tree.Name,
					ParentID: directory_tree.ParentID,
				})
				tree = append(tree, directory_tree.ParentID)
				node = directory_tree.ParentID
			}

			if directory_tree.Id == node && directory_tree.ParentID == 0 {
				nodes = append(nodes, dto.GetDirectoryTreeRes{
					Id:       directory_tree.Id,
					Name:     directory_tree.Name,
					ParentID: directory_tree.ParentID,
				})
				node = 0
			}
		}
	}
	for i, j := 0, len(tree)-1; i < j; i, j = i+1, j-1 {
		tree[i], tree[j] = tree[j], tree[i]
	}
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}

	return &tree, &nodes, nil
}

func (s *serviceDirectoryTree) ListByChild(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	directory, err := s.ListByParent(req)
	if err != nil {
		b, err := s.directoryTreeRepository.GetDirectoryParent(req)

		if err != nil {
			return nil, err
		}
		return b, err
	}
	if directory.ParentID == 0 {
		return directory, nil
	}
	route, _, err := s.GetRoute(req)

	if err != nil {
		return nil, err
	}
	tree := *route

	node_root, err := s.ListByParent(&dto.GetDirectoryTreeReq{Id: tree[0], InstitutionID: req.InstitutionID, AsesorID: req.AsesorID})
	if err != nil {
		return nil, err
	}
	var node_child *dto.GetDirectoryTreeRes
	node_child = node_root

	for i, j := 1, len(tree); i < j; i = i + 1 {
		directories := node_child.Directory
		for _, directory_tree := range directories {
			if directory_tree.Id == tree[i] {
				directory_child, _ := s.ListByParent(&dto.GetDirectoryTreeReq{Id: directory_tree.Id, InstitutionID: req.InstitutionID, AsesorID: req.AsesorID})
				parent := directory_tree
				parent.Open = true
				parent.Directory = directory_child.Directory
				parent.Document = directory_child.Document
				node_child = directory_tree
				break
			}
		}
	}
	return node_root, nil
}

func (s *serviceDirectoryTree) Delete(req *dto.GetDirectoryTreeReq) (bool, error) {
	repoReq := &sq.DeleteDirectoryTreeByIdParams{
		DirectoryID: req.Id,
		DeletedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	res_parent, err := s.ListByParent(req)
	if err != nil {
		return false, err
	}
	if len(res_parent.Directory) > 0 || len(*res_parent.Document) > 0 {
		return false, errors.New("Directory has files")
	}

	if err := s.directoryTreeRepository.DeleteDirectory(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}

func (s *serviceDirectoryTree) Update(req *dto.UpdateDirectoryTreeReq) (bool, error) {
	directory, err := s.Get(&dto.GetDirectoryTreeReq{Id: req.DirectoryId, InstitutionID: req.InstitutionID})

	if err != nil {
		return false, err
	}

	repoReq := &sq.UpdateDirectoryTreeByIdParams{
		DirectoryID: req.DirectoryId,
		Name:        pgtype.Text{String: directory.Name, Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		ParentID:    pgtype.Int8{Int64: directory.ParentID, Valid: true},
		AsesorID:    pgtype.Int4{Int32: directory.AsesorID, Valid: true},
	}

	if req.Name != "" {
		repoReq.Name = pgtype.Text{String: req.Name, Valid: true}
	}

	if req.ParentID != 0 {
		repoReq.ParentID = pgtype.Int8{Int64: req.ParentID, Valid: true}
	} else {
		repoReq.ParentID = pgtype.Int8{Int64: directory.ParentID, Valid: true}
	}

	if req.AsesorID != 0 {
		repoReq.AsesorID = pgtype.Int4{Int32: directory.AsesorID, Valid: true}
	}

	err = s.directoryTreeRepository.UpdateDirectoryTree(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceDirectoryTree) GetIdByParent(req *dto.GetDirectoryTreeReq) (*int64, error) {
	return s.directoryTreeRepository.GetDirectoryIdByParent(req)
}
